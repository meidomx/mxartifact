package gorepo

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/meidomx/mxartifact/config"

	"github.com/go-resty/resty/v2"
)

var _ GoModuleRepository = new(proxyRepo)

var defaultUpstreamRepo = "https://proxy.golang.org"
var defaultUpstreamSumdb = "https://sum.golang.org"
var altUpstreamSumdbForCn = "https://sum.golang.google.cn"

type proxyRepo struct {
	Name         string
	UpstreamRepo string

	HttpProxy string

	debug  bool
	client *resty.Client
}

func NewProxyRepo(conf config.GoRepoConf, gc *config.Config) GoModuleRepository {
	client := resty.New()

	if len(conf.HttpProxy) > 0 {
		client.SetProxy(conf.HttpProxy)
	}

	// initialize repo
	return &proxyRepo{
		Name:         conf.Name,
		UpstreamRepo: conf.UpstreamRepository,
		HttpProxy:    conf.HttpProxy,
		debug:        gc.Shared.Debug,
		client:       client,
	}
}

func (p *proxyRepo) fullUpstreamUri(uri string) string {
	if strings.HasPrefix(uri, "/sumdb/") {
		// sumdb request
		if len(p.UpstreamRepo) == 0 {
			// trim prefix
			uri = strings.TrimPrefix(uri, "/sumdb/")
			return fmt.Sprint("https://" + uri)
		} else {
			return fmt.Sprint(p.UpstreamRepo + uri)
		}
	} else {
		// proxy
		if len(p.UpstreamRepo) == 0 {
			return fmt.Sprint(defaultUpstreamRepo + uri)
		} else {
			return fmt.Sprint(p.UpstreamRepo + uri)
		}
	}
}

func (p *proxyRepo) decideErr(status int) error {
	if 200 <= status && status <= 299 {
		return nil
	} else if status == 404 {
		return ErrResourceNotFound
	} else if status == 403 {
		return ErrResourceForbidden
	} else if 400 <= status && status <= 499 {
		return ErrResourceClientError
	} else if 500 <= status && status <= 599 {
		return ErrResourceServerError
	}
	return errors.New("unknown http status:" + fmt.Sprint(status))
}

func (p *proxyRepo) FetchResource(uri string) ([]byte, string, error) {
	resourceUri := p.fullUpstreamUri(uri)
	if p.debug {
		log.Println("Fetch remote resource on:" + resourceUri)
	}
	res, err := p.client.R().Get(resourceUri)
	if err != nil {
		return nil, "", err
	}
	if err := p.decideErr(res.StatusCode()); err != nil {
		return nil, "", err
	}
	return res.Body(), res.Header().Get("Content-Type"), nil
}

func (p *proxyRepo) SupportSumDBProxy(uri string) (bool, error) {
	if !strings.HasSuffix(uri, "/supported") {
		return false, errors.New("not check sumdb proxy uri")
	}
	if len(p.UpstreamRepo) > 0 {
		// upstream repo mode
		_, _, err := p.FetchResource(uri)
		if err != nil {
			log.Println("sumdb unsupported reason:" + fmt.Sprint(err))
			return false, nil
		} else {
			return true, nil
		}
	}
	return true, nil
}

func (p *proxyRepo) UploadResource(uri string, data []byte) error {
	return errors.New("upload unsupported")
}

func (p *proxyRepo) SupportUpload(uri string) (bool, error) {
	return false, nil
}
