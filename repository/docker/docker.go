package docker

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/meidomx/mxartifact/config"
	"github.com/meidomx/mxartifact/resource"
)

const (
	BindOptionHostName = "hostname"
)

var (
	end2Hash *regexp.Regexp
	end3Hash *regexp.Regexp
	end3Tag  *regexp.Regexp
)

func init() {
	var err error
	if end2Hash, err = regexp.Compile(`/blobs/[a-zA-Z0-9_+.-]+:[a-zA-Z0-9=_-]+$`); err != nil {
		panic(err)
	}
	if end3Hash, err = regexp.Compile(`/manifests/[a-zA-Z0-9_+.-]+:[a-zA-Z0-9=_-]+$`); err != nil {
		panic(err)
	}
	if end3Tag, err = regexp.Compile(`/manifests/[a-zA-Z0-9_][a-zA-Z0-9._-]{0,127}$`); err != nil {
		panic(err)
	}
}

func Init(engine *resource.ResourceManager, c *config.Config) {
	for _, docker := range c.Repository.Dockers {
		//FIXME support all types of docker repo
		if docker.Type == "local" {
			log.Println("[WARN] docker repository: unsupported type: [" + docker.Type + "] name: " + docker.Name)
			continue
		} else if docker.Type == "proxy" {
			// initialize docker repository
			var repo DockerRepository = NewProxiedDockerRepository(docker.HttpProxy, docker.UpstreamRepository)

			// register listener
			if len(docker.BindListeners) > 0 {
				for _, bind := range docker.BindListeners {
					// add char / to the end of baseUrl
					baseUrl := docker.BaseUrl
					if !strings.HasSuffix(baseUrl, "/") {
						baseUrl += "/"
					}
					// register http handler
					engine.AddHttpResourceConsumer(bind.Name, bind.Options[BindOptionHostName], baseUrl, func(w http.ResponseWriter, r *http.Request) {
						log.Println("handle docker request:", r.Method, r.URL.Path)

						// strip requestURI by baseURL
						dockerUri, _ := strings.CutPrefix(r.URL.Path, baseUrl)
						if !strings.HasPrefix(dockerUri, "/") {
							dockerUri = "/" + dockerUri
						}

						// only support v2 apis
						if !strings.HasPrefix(dockerUri, "/v2/") {
							http.NotFound(w, r)
							return
						}
						// api: end-1
						if dockerUri == "/v2/" {
							w.WriteHeader(http.StatusOK)
							return
						}
						// pull image operations
						if r.Method == "GET" || r.Method == "HEAD" {
							log.Println("pull image operations:", r.Method, dockerUri)
							// end2
							// e.g. http://127.0.0.1/v2/myorg/myrepo/blobs/sha256:b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
							if end2Hash.MatchString(dockerUri) {
								repo.HandlerRetrieveBlob(r.Method, dockerUri)(w, r)
								return
							}
							// end3 tag
							// e.g. http://127.0.0.1/v2/myorg/myrepo/manifests/tagtest0
							if end3Tag.MatchString(dockerUri) {
								repo.HandlerRetrieveManifestByTag(r.Method, dockerUri)(w, r)
								return
							}
							// end3 hash
							// e.g. http://127.0.0.1/v2/myorg/myrepo/manifests/sha256:9dad4f699e3a49f674aeec428e0a973b8101800866569a3a9b96647c7a4fd238
							if end3Hash.MatchString(dockerUri) {
								repo.HandlerRetrieveManifestByHash(r.Method, dockerUri)(w, r)
								return
							}
						}
						// default: 404
						http.NotFound(w, r)
						return
					})
				}
			}
		}
	}
}
