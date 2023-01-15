# mxartifact

An opensource repository service.

## 1. Features

* [x] Repository: Go Module
* [x] Repository - Go: upstream type: proxy server
* [x] Repository - Go: upstream type: L3/L4 proxy server
* [x] Repository - Go: local storage
* [ ] Upstream: http proxy support

### 1.1 Golang

#### Prerequisite

* Install `golang >= 1.19`

## 2. TODO List

* [ ] Separate urls for different repositories
* [ ] Repository: Maven
* [ ] Repository: Nuget
* [ ] Repository: Cargo
* [ ] Repository: Docker
* [ ] Private repository: layered(nested and recursive) repositories, file permission(rw) on repository
* [ ] Repository: standard/proxy/layered repository
* [ ] Upstream: Queuing
* [ ] Repository - Go: need extract request type before processing repositories & rely on the type
* [ ] Repository - Go: support safe concurrent local storage
* [ ] Repository - Go: simulating GOPRIVATE
* [ ] Repository - Go: upstream - A.source repository C.internal recursive repository
* [ ] Repository - Go: checksum database
* [ ] Repository - Go: Max file size defined in the Go Module specification
* [ ] Repository - Go: support management for local persisted repositories including uploading/deleting files
* [ ] Workflow support: Production/Pre published/Staged/Develop/Multiple environment isolation
* [ ] Server: cluster and HA
* [ ] Security: ratelimiter

## 3. Dependency

* [x] https://github.com/goproxy/goproxy - go module proxy
* [x] https://github.com/spf13/afero - file system support
* [x] https://github.com/BurntSushi/toml - toml config file parser
* [x] https://github.com/go-resty/resty - http client

Reference implementations

##### Go

* https://github.com/goproxy/goproxy
* https://github.com/goproxyio/goproxy

## 4. Client support

#### 4.1 Go

* Sumdb: to avoid sumdb check for a private repository, use GONOSUMDB env. GOPRIVATE is an alternative if private proxy
  function is not enabled and used.

## A. Reference documents

### Go

[Go Module](https://go.dev/ref/mod)
[Go Module Sumdb Proxy](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database)
