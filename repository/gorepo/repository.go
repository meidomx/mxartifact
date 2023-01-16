package gorepo

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/meidomx/mxartifact/config"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine, c *config.Config) {

	repoMap := make(map[string]GoModuleRepository)

	for _, v := range c.Repository.Golangs {
		log.Printf("start golang repo [%s] on [%s]", v.Name, v.BaseUrl)

		var client GoModuleRepository
		switch v.Type {
		case "proxy":
			client = NewProxyRepo(v, c)
		case "local":
			pRepo, ok := repoMap[v.ParentRepository]
			if !ok {
				panic(errors.New("unknown parent repository:" + fmt.Sprint(v.ParentRepository)))
			}
			client = NewLocalCacheRepo(v, c, pRepo)
		default:
			panic("unknown repository type:" + v.Type)
		}

		if _, ok := repoMap[v.Name]; ok {
			panic(errors.New("already have the GoModuleRepository with name:" + v.Name))
		}
		repoMap[v.Name] = client

		if len(v.BaseUrl) > 0 {
			group := engine.Group(v.BaseUrl)
			group.GET("/*goquery", func(context *gin.Context) {
				if c.Shared.Debug {
					log.Println("go request:" + context.Param("goquery"))
				}

				// block access unknown uri including / , /favicon.ico , etc
				pathType := extractPathType(context.Param("goquery"))
				if len(pathType) <= 0 {
					if c.Shared.Debug {
						log.Println("unknown pathType for:" + context.Param("goquery"))
					}
					context.Status(http.StatusNotFound)
					return
				}

				data, ct, err := client.FetchResource(context.Param("goquery"), pathType)
				if err == ErrResourceNotFound {
					context.Status(http.StatusNotFound)
					return
				} else if err == ErrResourceForbidden {
					context.Status(http.StatusForbidden)
					return
				} else if err == ErrResourceClientError {
					context.Status(http.StatusBadRequest)
					return
				} else if err == ErrResourceServerError || err != nil {
					context.Status(http.StatusInternalServerError)
					return
				} else {
					context.Data(http.StatusOK, ct, data)
					return
				}
			})
		}
	}

}

func extractPathType(param string) string {
	if strings.HasSuffix(param, "/@v/list") {
		return PathTypeListModule
	} else if strings.HasSuffix(param, ".info") {
		return PathTypeModuleMeta
	} else if strings.HasSuffix(param, ".mod") {
		return PathTypeModuleVersion
	} else if strings.HasSuffix(param, ".zip") {
		return PathTypeModuleContent
	} else if strings.HasSuffix(param, "/@latest") {
		return PathTypeModuleLatestVersion
	} else if strings.HasPrefix(param, "/sumdb/") {
		if strings.HasSuffix(param, "/supported") {
			return PathTypeSumProxySupported
		} else {
			return PathTypeSumProxy
		}
	} else {
		// unknown type
		return ""
	}
}

const (
	PathTypeListModule          = "$base/$module/@v/list"
	PathTypeModuleMeta          = "$base/$module/@v/$version.info"
	PathTypeModuleVersion       = "$base/$module/@v/$version.mod"
	PathTypeModuleContent       = "$base/$module/@v/$version.zip"
	PathTypeModuleLatestVersion = "$base/$module/@latest"

	PathTypeSumProxy          = "<proxyURL>/sumdb/<databaseURL>"
	PathTypeSumProxySupported = "<proxyURL>/sumdb/<sumdb-name>/supported"
)

const (
	TriggerLoadAllVersionsAndLatestVersion = "PathTypeListModule,PathTypeModuleLatestVersion"
)

var (
	ErrResourceNotFound    = errors.New("not found/404")
	ErrResourceForbidden   = errors.New("forbidden/403")
	ErrResourceClientError = errors.New("other client error/4xx")
	ErrResourceServerError = errors.New("server error/5xx")
)

type GoModuleRepository interface {
	// FetchResource fetch uri data & content-type with error if failed
	FetchResource(uri string, pathType string) ([]byte, string, error)

	SupportSumDBProxy(uri string) (bool, error)

	UploadResource(uri string, data []byte) error
	SupportUpload(uri string) (bool, error)
}
