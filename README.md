# mxartifact

An opensource repository service.

## 1. Features

* [ ] Golang: go module

### 1.1 Golang

#### Prerequisite

* Install `golang >= 1.19`

#### Features

* [x] Repository: Go Module
* [x] Repository - Go: upstream type: proxy server
* [x] Repository - Go: upstream type: L3/L4 proxy server
* [ ] Upstream: http proxy support

#### TODO List

* [ ] Separate urls for different repositories
* [ ] Repository: Maven
* [ ] Repository: Nuget
* [ ] Repository: Cargo
* [ ] Repository: Docker
* [ ] Private repository: layered(nested and recursive) repositories, file permission(rw) on repository
* [ ] Repository: standard/proxy/layered repository
* [ ] Upstream: Queuing
* [ ] Repository - Go: simulating GOPRIVATE
* [ ] Repository - Go: upstream - A.source repository C.internal recursive repository D.local storage
* [ ] Repository - Go: checksum database
* [ ] Repository - Go: upstream type: proxy server
* [ ] Repository - Go: upstream type: L3/L4 proxy server
* [ ] Repository - Go: Max file size defined in the Go Module specification
* [ ] Workflow support: Production/Pre published/Staged/Develop/Multiple environment isolation
* [ ] Server: cluster and HA

## 2. Dependency

* [x] https://github.com/goproxy/goproxy - go module proxy
* [x] https://github.com/spf13/afero - file system support
* [x] https://github.com/BurntSushi/toml - toml config file parser

## 3. Client support

#### 3.1 Go

* Sumdb: to avoid sumdb check for a private repository, use GONOSUMDB env. GOPRIVATE is an alternative if private proxy function is not enabled and used.

## A. Reference documents

### Go

[Go Module](https://go.dev/ref/mod)
[Go Module Sumdb Proxy](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database)
