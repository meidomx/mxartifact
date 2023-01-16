package gorepo

import (
	"fmt"
	"log"
	"strings"

	"github.com/meidomx/mxartifact/config"

	"github.com/spf13/afero"
)

type localCacheRepo struct {
	fs         afero.Fs
	sourceRepo GoModuleRepository
}

func NewLocalCacheRepo(conf config.GoRepoConf, c *config.Config, sourceRepo GoModuleRepository) GoModuleRepository {
	subpath := fmt.Sprint(c.FileStorage.Location, afero.FilePathSeparator, conf.Name)
	log.Printf("prepare folder [%s] for go repo [%s]", subpath, conf.Name)
	osfs := afero.NewOsFs()
	if ok, err := afero.Exists(osfs, subpath); err != nil {
		panic(err)
	} else if !ok {
		if err := osfs.MkdirAll(subpath, 0755); err != nil {
			panic(err)
		}
	}
	fs := afero.NewBasePathFs(osfs, subpath)

	return &localCacheRepo{
		fs:         fs,
		sourceRepo: sourceRepo,
	}
}

func (l *localCacheRepo) FetchResource(uri string, pathType string) ([]byte, string, error) {
	// find local cache
	if ok, err := afero.Exists(l.fs, uri); err != nil {
		return nil, "", err
	} else if ok {
		if d, err := afero.ReadFile(l.fs, uri); err != nil {
			return nil, "", err
		} else {
			var contentType string
			if strings.HasSuffix(strings.ToLower(uri), ".zip") {
				contentType = "application/zip"
			} else if strings.HasSuffix(strings.ToLower(uri), ".info") {
				contentType = "application/json"
			} else {
				contentType = "text/plain; charset=utf-8"
			}
			return d, contentType, nil
		}
	}

	d, ct, err := l.sourceRepo.FetchResource(uri, pathType)
	if err != nil {
		return nil, "", err
	}

	//FIXME should not record unexpected file
	// cache file
	folderPath := uri[0:strings.LastIndex(uri, "/")]
	if err := l.fs.MkdirAll(folderPath, 0755); err != nil {
		return nil, "", err
	}
	if err := afero.WriteFile(l.fs, uri, d, 0644); err != nil {
		return nil, "", err
	}

	return d, ct, nil
}

func (l *localCacheRepo) SupportSumDBProxy(uri string) (bool, error) {
	return l.sourceRepo.SupportSumDBProxy(uri)
}

func (l *localCacheRepo) UploadResource(uri string, data []byte) error {
	return l.sourceRepo.UploadResource(uri, data)
}

func (l *localCacheRepo) SupportUpload(uri string) (bool, error) {
	return l.sourceRepo.SupportUpload(uri)
}

var _ GoModuleRepository = new(localCacheRepo)
