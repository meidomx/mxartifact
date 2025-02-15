package shared

import (
	"io/fs"
	"os"
)

type FS interface {
	fs.FS
	fs.ReadDirFS
	fs.StatFS
	fs.ReadFileFS
}

func Exists(fs FS, path string) (bool, error) {
	_, err := fs.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	} else {
		return false, err
	}
}
