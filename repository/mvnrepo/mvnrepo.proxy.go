package mvnrepo

import (
	"errors"
	"fmt"
	"log"

	"github.com/meidomx/mxartifact/config"

	"github.com/go-resty/resty/v2"
)

const (
	MavenCentralRepoUrl = "https://repo1.maven.org/maven2"
)

type mvnProxyRepo struct {
	Name         string
	UpstreamRepo string
	HttpProxy    string

	client *resty.Client
	debug  bool
}

func NewMvnProxyRepo(conf config.MavenRepoConf, c *config.Config) MavenRepository {
	client := resty.New()

	//FIXME should not use SetTimeout since this method limits the total time for the whole request
	// This value may not be enough for huge files.
	// Need find another way to detect and cut down conn/read timeout connection
	//client.SetTimeout(10 * time.Second)

	if len(conf.HttpProxy) > 0 {
		client.SetProxy(conf.HttpProxy)
	}

	// initialize repo
	return &mvnProxyRepo{
		Name:         conf.Name,
		UpstreamRepo: conf.UpstreamRepository,
		HttpProxy:    conf.HttpProxy,
		debug:        c.Shared.Debug,
		client:       client,
	}
}

func (m *mvnProxyRepo) FetchResource(uri string) ([]byte, string, error) {
	resourceUri := m.fullUpstreamUri(uri)
	if m.debug {
		log.Println("Fetch remote resource on: " + resourceUri + " with http_proxy=" + fmt.Sprint(len(m.HttpProxy) > 0))
	}
	res, err := m.client.R().Get(resourceUri)
	if err != nil {
		return nil, "", err
	}
	if err := m.decideErr(res.StatusCode()); err != nil {
		return nil, "", err
	}
	return res.Body(), res.Header().Get("Content-Type"), nil
}

func (m *mvnProxyRepo) MetaResource(uri string) (string, error) {
	resourceUri := m.fullUpstreamUri(uri)
	if m.debug {
		log.Println("Fetch meta from remote resource on: " + resourceUri + " with http_proxy=" + fmt.Sprint(len(m.HttpProxy) > 0))
	}
	res, err := m.client.R().Head(resourceUri)
	if err != nil {
		return "", err
	}
	if err := m.decideErr(res.StatusCode()); err != nil {
		return "", err
	}
	return res.Header().Get("Content-Type"), nil
}

func (m *mvnProxyRepo) UploadResource(uri string, data []byte) error {
	return ErrNotSupported
}

func (m *mvnProxyRepo) RemoveResource(uri string) error {
	return ErrNotSupported
}

func (m *mvnProxyRepo) fullUpstreamUri(uri string) string {
	var baseUrl string
	if len(m.UpstreamRepo) > 0 {
		baseUrl = m.UpstreamRepo
	} else {
		baseUrl = MavenCentralRepoUrl
	}

	return baseUrl + uri
}

func (m *mvnProxyRepo) decideErr(status int) error {
	// for 3xx codes, they should be automatically followed by http client and well taken care
	if 200 <= status && status <= 299 {
		return nil
	} else if status == 404 {
		return ErrResourceNotFound
	} else if status == 403 {
		return ErrResourceForbidden
	} else if 400 <= status && status <= 499 {
		return ErrOperationFailed
	} else if 500 <= status && status <= 599 {
		return ErrOperationFailed
	}
	return errors.New("unknown http status:" + fmt.Sprint(status))
}

var _ MavenRepository = new(mvnProxyRepo)
