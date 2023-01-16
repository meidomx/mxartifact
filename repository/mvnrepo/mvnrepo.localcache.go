package mvnrepo

import (
	"fmt"
	"log"
	"strings"

	"github.com/meidomx/mxartifact/config"

	"github.com/spf13/afero"
)

type mvnLocalCacheRepo struct {
	fs         afero.Fs
	sourceRepo MavenRepository
}

func NewMvnLocalCacheRepo(conf config.MavenRepoConf, c *config.Config, sourceRepo MavenRepository) MavenRepository {
	subpath := fmt.Sprint(c.FileStorage.Location, afero.FilePathSeparator, conf.Name)
	log.Printf("prepare folder [%s] for mvn repo [%s]", subpath, conf.Name)
	osfs := afero.NewOsFs()
	if ok, err := afero.Exists(osfs, subpath); err != nil {
		panic(err)
	} else if !ok {
		if err := osfs.MkdirAll(subpath, 0755); err != nil {
			panic(err)
		}
	}
	fs := afero.NewBasePathFs(osfs, subpath)

	return &mvnLocalCacheRepo{
		fs:         fs,
		sourceRepo: sourceRepo,
	}
}

func (m *mvnLocalCacheRepo) FetchResource(uri string) ([]byte, string, error) {
	// find local cache
	if ok, err := afero.Exists(m.fs, uri); err != nil {
		return nil, "", err
	} else if ok {
		if d, err := afero.ReadFile(m.fs, uri); err != nil {
			return nil, "", err
		} else {
			var contentType string
			if strings.HasSuffix(strings.ToLower(uri), ".pom") {
				contentType = ContentTypeFilePom
			} else if strings.HasSuffix(strings.ToLower(uri), ".jar") {
				contentType = ContentTypeFileJar
			} else if strings.HasSuffix(strings.ToLower(uri), ".xml") {
				contentType = ContentTypeFileXml
			} else {
				contentType = ContentTypeFileOther
			}
			return d, contentType, nil
		}
	}

	d, ct, err := m.sourceRepo.FetchResource(uri)
	if err != nil {
		return nil, "", err
	}

	//FIXME should not record unexpected file
	// cache file
	folderPath := uri[0:strings.LastIndex(uri, "/")]
	if err := m.fs.MkdirAll(folderPath, 0755); err != nil {
		return nil, "", err
	}
	if err := afero.WriteFile(m.fs, uri, d, 0644); err != nil {
		return nil, "", err
	}

	return d, ct, nil
}

func (m *mvnLocalCacheRepo) MetaResource(uri string) (string, error) {
	if ok, err := afero.Exists(m.fs, uri); err != nil {
		return "", err
	} else if ok {
		var contentType string
		if strings.HasSuffix(strings.ToLower(uri), ".pom") {
			contentType = ContentTypeFilePom
		} else if strings.HasSuffix(strings.ToLower(uri), ".jar") {
			contentType = ContentTypeFileJar
		} else if strings.HasSuffix(strings.ToLower(uri), ".xml") {
			contentType = ContentTypeFileXml
		} else {
			contentType = ContentTypeFileOther
		}
		return contentType, nil
	}

	return m.sourceRepo.MetaResource(uri)
}

func (m *mvnLocalCacheRepo) UploadResource(uri string, data []byte) error {
	return ErrNotSupported
}

func (m *mvnLocalCacheRepo) RemoveResource(uri string) error {
	return ErrNotSupported
}

var _ MavenRepository = new(mvnLocalCacheRepo)
