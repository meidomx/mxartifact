package mvnrepo

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/meidomx/mxartifact/config"

	"github.com/gin-gonic/gin"
)

type MavenRepository interface {
	FetchResource(uri string) ([]byte, string, error)
	MetaResource(uri string) (string, error)

	UploadResource(uri string, data []byte) error
	RemoveResource(uri string) error
}

var (
	ErrResourceNotFound  = errors.New("resource not found")
	ErrResourceForbidden = errors.New("forbidden")
	ErrOperationFailed   = errors.New("operation failed")
	ErrNotSupported      = errors.New("not supported")
)

const (
	ContentTypeFilePom   = "text/xml"
	ContentTypeFileXml   = "text/xml"
	ContentTypeFileJar   = "application/java-archive"
	ContentTypeFileOther = "text/plain"
)

func Init(engine *gin.Engine, c *config.Config) {

	repoMap := make(map[string]MavenRepository)

	for _, v := range c.Repository.Mavens {
		log.Printf("start maven repo [%s] on [%s]", v.Name, v.BaseUrl)

		var client MavenRepository
		switch v.Type {
		case "proxy":
			client = NewMvnProxyRepo(v, c)
		case "local":
			pRepo, ok := repoMap[v.ParentRepository]
			if !ok {
				panic(errors.New("unknown parent repository:" + fmt.Sprint(v.ParentRepository)))
			}
			client = NewMvnLocalCacheRepo(v, c, pRepo)
		default:
			panic("unknown repository type:" + v.Type)
		}

		if _, ok := repoMap[v.Name]; ok {
			panic(errors.New("already have the MavenRepository with name:" + v.Name))
		}
		repoMap[v.Name] = client

		if len(v.BaseUrl) > 0 {
			group := engine.Group(v.BaseUrl)
			group.GET("/*query", func(context *gin.Context) {
				if c.Shared.Debug {
					log.Println("maven request:" + context.Param("query"))
				}

				data, ct, err := client.FetchResource(context.Param("query"))
				if err == ErrResourceNotFound {
					context.Status(http.StatusNotFound)
					return
				} else if err == ErrResourceForbidden {
					context.Status(http.StatusForbidden)
					return
				} else if err == ErrOperationFailed || err != nil {
					context.Status(http.StatusInternalServerError)
					return
				} else {
					context.Data(http.StatusOK, ct, data)
					return
				}
			})
			group.HEAD("/*query", func(context *gin.Context) {
				if c.Shared.Debug {
					log.Println("maven request:" + context.Param("query"))
				}

				ct, err := client.MetaResource(context.Param("query"))
				if err == ErrResourceNotFound {
					context.Status(http.StatusNotFound)
					return
				} else if err == ErrResourceForbidden {
					context.Status(http.StatusForbidden)
					return
				} else if err == ErrOperationFailed || err != nil {
					context.Status(http.StatusInternalServerError)
					return
				} else {
					context.Data(http.StatusOK, ct, nil)
					return
				}
			})
		}
	}

}
