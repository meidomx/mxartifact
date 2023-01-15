package main

import (
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/meidomx/mxartifact/config"
	"github.com/meidomx/mxartifact/repository/gorepo"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
)

func main() {
	fs := afero.NewOsFs()
	f, err := fs.Open("config.toml.example")
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

	r := gin.Default()
	gorepo.Init(r, cfg)
	log.Println("starting service on :9051 ...")
	if err := r.Run(":9051"); err != nil {
		log.Fatalln("start failed:" + fmt.Sprint(err))
	}
}
