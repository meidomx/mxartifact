package gorepo

import (
	"errors"
	"log"
	"net/http"

	"github.com/meidomx/mxartifact/config"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine, c *config.Config) {

	for _, v := range c.Repository.Golangs {
		log.Printf("start golang repo [%s] on [%s]", v.Name, v.BaseUrl)
		group := engine.Group(v.BaseUrl)

		var client GoModuleRepository
		switch v.Type {
		case "proxy":
			client = NewProxyRepo(v, c)
		default:
			panic("unknown repository type:" + v.Type)
		}

		group.GET("/*goquery", func(context *gin.Context) {
			if c.Shared.Debug {
				log.Println("go request:" + context.Param("goquery"))
			}

			//FIXME block access unknown uri including / , /favicon.ico , etc

			data, ct, err := client.FetchResource(context.Param("goquery"))
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
	FetchResource(uri string) ([]byte, string, error)

	SupportSumDBProxy(uri string) (bool, error)

	UploadResource(uri string, data []byte) error
	SupportUpload(uri string) (bool, error)
}
