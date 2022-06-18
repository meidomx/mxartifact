package main

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/meidomx/config"
	"github.com/meidomx/repository/golang"
	"io"

	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()
	f, err := fs.Open("config.toml")
	if err != nil {
		panic(errors.New(fmt.Sprint("open config.toml file error:", err)))
	}
	defer f.Close()
	configData, err := io.ReadAll(f)
	if err != nil {
		panic(errors.New(fmt.Sprint("read config.toml file error:", err)))
	}
	cfg := new(config.Config)
	if err := toml.Unmarshal(configData, cfg); err != nil {
		panic(errors.New(fmt.Sprint("parse config error:", err)))
	}
	var _ = f.Close()

	golang.Init()
}
