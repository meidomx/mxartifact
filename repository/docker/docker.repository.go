package docker

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/meidomx/mxartifact/shared"
)

type DockerRepository interface {
	//HandlerRetrieveBlob implements endpoint: end-2
	HandlerRetrieveBlob(method, url string) shared.HttpHandler
	//HandlerRetrieveManifestByTag implements endpoint: end-3 with tag as reference
	HandlerRetrieveManifestByTag(method, url string) shared.HttpHandler
	//HandlerRetrieveManifestByHash implements endpoint: end-3 with hash as reference
	HandlerRetrieveManifestByHash(method, url string) shared.HttpHandler
}

type ProxiedDockerRepository struct {
	client   *resty.Client
	upstream string
}

func (p *ProxiedDockerRepository) HandlerRetrieveBlob(method, url string) shared.HttpHandler {
	if method == http.MethodHead {
		return func(w http.ResponseWriter, r *http.Request) {
			remote := fmt.Sprintf("%s%s", p.upstream, url)
			accepts := r.Header.Values("Accept")
			log.Println("docker proxy remote head:", remote, len(accepts), accepts)
			res, err := p.client.R().SetHeaderMultiValues(map[string][]string{
				"Accept": accepts,
			}).SetDoNotParseResponse(true).Head(remote)
			if err != nil {
				log.Println("HandlerRetrieveBlob failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer func(body io.ReadCloser) {
				_ = body.Close()
			}(res.RawBody())
			if res.StatusCode() == 200 {
				if ct := res.Header().Get("Content-Type"); ct != "" {
					w.Header().Set("Content-Type", ct)
				}
				if digest := res.Header().Get("Docker-Content-Digest"); digest != "" {
					w.Header().Set("Docker-Content-Digest", digest)
				}
				w.WriteHeader(res.StatusCode())
				return
			} else {
				w.WriteHeader(res.StatusCode())
				return
			}
		}
	} else if method == http.MethodGet {
		return func(w http.ResponseWriter, r *http.Request) {
			remote := fmt.Sprintf("%s%s", p.upstream, url)
			accepts := r.Header.Values("Accept")
			log.Println("docker proxy remote get:", remote, len(accepts), accepts)
			res, err := p.client.R().SetHeaderMultiValues(map[string][]string{
				"Accept": accepts,
			}).SetDoNotParseResponse(true).Get(remote)
			if err != nil {
				log.Println("HandlerRetrieveBlob failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer func(body io.ReadCloser) {
				_ = body.Close()
			}(res.RawBody())
			if res.StatusCode() == 200 {
				var ct, digest, cl string
				if ct = res.Header().Get("Content-Type"); ct != "" {
					w.Header().Set("Content-Type", ct)
				}
				if digest = res.Header().Get("Docker-Content-Digest"); digest != "" {
					w.Header().Set("Docker-Content-Digest", digest)
				}
				if cl = res.Header().Get("Content-Length"); cl != "" {
					w.Header().Set("Content-Length", cl)
				}
				log.Printf("HandlerRetrieveBlob downloading: [%s] [%s] [%s] [%s]\n", remote, ct, digest, cl)
				w.WriteHeader(res.StatusCode())
				if _, err := io.Copy(w, res.RawBody()); err != nil {
					log.Println("HandlerRetrieveBlob copy body failed:", err)
				}
				return
			} else {
				w.WriteHeader(res.StatusCode())
				return
			}
		}
	} else {
		panic(errors.New("method not supported:" + method))
	}
}

func (p *ProxiedDockerRepository) HandlerRetrieveManifestByTag(method, url string) shared.HttpHandler {
	if method == http.MethodHead {
		return func(w http.ResponseWriter, r *http.Request) {
			remote := fmt.Sprintf("%s%s", p.upstream, url)
			accepts := r.Header.Values("Accept")
			log.Println("docker proxy remote head:", remote, len(accepts), accepts)
			res, err := p.client.R().SetHeaderMultiValues(map[string][]string{
				"Accept": accepts,
			}).SetDoNotParseResponse(true).Head(remote)
			if err != nil {
				log.Println("HandlerRetrieveManifestByTag failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer func(body io.ReadCloser) {
				_ = body.Close()
			}(res.RawBody())
			if res.StatusCode() == 200 {
				if ct := res.Header().Get("Content-Type"); ct != "" {
					w.Header().Set("Content-Type", ct)
				}
				if digest := res.Header().Get("Docker-Content-Digest"); digest != "" {
					w.Header().Set("Docker-Content-Digest", digest)
				}
				w.WriteHeader(res.StatusCode())
				return
			} else {
				w.WriteHeader(res.StatusCode())
				return
			}
		}
	} else if method == http.MethodGet {
		return func(w http.ResponseWriter, r *http.Request) {
			remote := fmt.Sprintf("%s%s", p.upstream, url)
			accepts := r.Header.Values("Accept")
			log.Println("docker proxy remote get:", remote, len(accepts), accepts)
			res, err := p.client.R().SetHeaderMultiValues(map[string][]string{
				"Accept": accepts,
			}).SetDoNotParseResponse(true).Get(remote)
			if err != nil {
				log.Println("HandlerRetrieveManifestByTag failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer func(body io.ReadCloser) {
				_ = body.Close()
			}(res.RawBody())
			if res.StatusCode() == 200 {
				var ct, digest, cl string
				if ct = res.Header().Get("Content-Type"); ct != "" {
					w.Header().Set("Content-Type", ct)
				}
				if digest = res.Header().Get("Docker-Content-Digest"); digest != "" {
					w.Header().Set("Docker-Content-Digest", digest)
				}
				if cl = res.Header().Get("Content-Length"); cl != "" {
					w.Header().Set("Content-Length", cl)
				}
				log.Printf("HandlerRetrieveManifestByTag downloading: [%s] [%s] [%s] [%s]\n", remote, ct, digest, cl)
				w.WriteHeader(res.StatusCode())
				if _, err := io.Copy(w, res.RawBody()); err != nil {
					log.Println("HandlerRetrieveManifestByTag copy body failed:", err)
				}
				return
			} else {
				w.WriteHeader(res.StatusCode())
				return
			}
		}
	} else {
		panic(errors.New("method not supported:" + method))
	}
}

func (p *ProxiedDockerRepository) HandlerRetrieveManifestByHash(method, url string) shared.HttpHandler {
	if method == http.MethodHead {
		return func(w http.ResponseWriter, r *http.Request) {
			remote := fmt.Sprintf("%s%s", p.upstream, url)
			accepts := r.Header.Values("Accept")
			log.Println("docker proxy remote head:", remote, len(accepts), accepts)
			res, err := p.client.R().SetHeaderMultiValues(map[string][]string{
				"Accept": accepts,
			}).SetDoNotParseResponse(true).Head(remote)
			if err != nil {
				log.Println("HandlerRetrieveManifestByHash failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer func(body io.ReadCloser) {
				_ = body.Close()
			}(res.RawBody())
			if res.StatusCode() == 200 {
				if ct := res.Header().Get("Content-Type"); ct != "" {
					w.Header().Set("Content-Type", ct)
				}
				if digest := res.Header().Get("Docker-Content-Digest"); digest != "" {
					w.Header().Set("Docker-Content-Digest", digest)
				}
				w.WriteHeader(res.StatusCode())
				return
			} else {
				w.WriteHeader(res.StatusCode())
				return
			}
		}
	} else if method == http.MethodGet {
		return func(w http.ResponseWriter, r *http.Request) {
			remote := fmt.Sprintf("%s%s", p.upstream, url)
			accepts := r.Header.Values("Accept")
			log.Println("docker proxy remote get:", remote, len(accepts), accepts)
			res, err := p.client.R().SetHeaderMultiValues(map[string][]string{
				"Accept": accepts,
			}).SetDoNotParseResponse(true).Get(remote)
			if err != nil {
				log.Println("HandlerRetrieveManifestByHash failed:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			defer func(body io.ReadCloser) {
				_ = body.Close()
			}(res.RawBody())
			if res.StatusCode() == 200 {
				var ct, digest, cl string
				if ct = res.Header().Get("Content-Type"); ct != "" {
					w.Header().Set("Content-Type", ct)
				}
				if digest = res.Header().Get("Docker-Content-Digest"); digest != "" {
					w.Header().Set("Docker-Content-Digest", digest)
				}
				if cl = res.Header().Get("Content-Length"); cl != "" {
					w.Header().Set("Content-Length", cl)
				}
				log.Printf("HandlerRetrieveManifestByHash downloading: [%s] [%s] [%s] [%s]\n", remote, ct, digest, cl)
				w.WriteHeader(res.StatusCode())
				if _, err := io.Copy(w, res.RawBody()); err != nil {
					log.Println("HandlerRetrieveManifestByHash copy body failed:", err)
				}
				return
			} else {
				w.WriteHeader(res.StatusCode())
				return
			}
		}
	} else {
		panic(errors.New("method not supported:" + method))
	}
}

func NewProxiedDockerRepository(httpProxy, upstream string) DockerRepository {
	client := resty.New()
	client.SetProxy(httpProxy)

	repo := &ProxiedDockerRepository{
		client:   client,
		upstream: upstream,
	}
	return repo
}
