package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/meidomx/mxartifact/config"
	"github.com/meidomx/mxartifact/repository/cargo"
	"github.com/meidomx/mxartifact/repository/docker"
	"github.com/meidomx/mxartifact/repository/gorepo"
	"github.com/meidomx/mxartifact/repository/mvnrepo"
	"github.com/meidomx/mxartifact/repository/nuget"
	"github.com/meidomx/mxartifact/resource"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
)

func main() {
	f, err := os.Open("config.toml")
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

	rmgr := resource.NewResourceManager(cfg)
	r := gin.Default()
	gorepo.Init(r, cfg)
	mvnrepo.Init(r, cfg)
	cargo.Init(r, cfg)
	nuget.Init(r, cfg)
	docker.Init(rmgr, cfg)
	log.Printf("starting service on %s ...", cfg.Shared.Listen)
	if err := rmgr.Startup(); err != nil {
		if err := rmgr.Shutdown(); err != nil {
			//FIXME log shutdown error
		}
		panic(err)
	}
	if err := r.Run(cfg.Shared.Listen); err != nil {
		log.Fatalln("start failed:" + fmt.Sprint(err))
	}
}
