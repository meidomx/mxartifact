# MxArtifact

An opensource repository service.

The purpose of the project is to provide a simple and customizable artifact repository, supporting modern languages.

## 1. Features

* [x] Repository: Go Module
* [x] Repository - Go: upstream type: proxy server
* [x] Repository - Go: upstream type: L3/L4 proxy server
* [x] Repository - Go: local cache
* [x] Repository: Maven
* [x] Repository - Maven: proxy as a mirror including central/self-hosted by upstream config
* [ ] Repository - Maven: deploy to repository - POST/PUT/DELETE
* [x] Upstream: http proxy support
* [x] Customizable Repository: layered(nested and recursive) repositories and file permission
* [x] Customizable Repository: separate url for each repository

### 1.1 Golang

#### Prerequisite

* Install `golang >= 1.19`

## 2. TODO List

* [ ] Repository: Nuget
* [ ] Repository: Cargo
* [ ] Repository: Docker
* [ ] Upstream: Queuing
* [ ] Repository - Maven: auth and permissions
* [ ] Repository - Maven: repository type - release / snapshot / etc.
* [ ] Repository - Maven: local storage
* [ ] Repository - Maven: deploy with auth - username & password
* [ ] Repository - Maven: metadata for deploying snapshots
* [ ] Repository - Go: need extract request type before processing repositories & rely on the type
* [ ] Repository - Go: support safe concurrent local storage
* [ ] Repository - Go: simulating GOPRIVATE
* [ ] Repository - Go: upstream - A.source repository C.internal recursive repository
* [ ] Repository - Go: checksum database
* [ ] Repository - Go: Max file size defined in the Go Module specification
* [ ] Repository - Go: support management for local persisted repositories including uploading/deleting files
* [ ] Repository - Go: access token for pulling private repository
* [ ] Repository: streaming transfer for large files in slow upstream env
* [ ] Workflow support: Production/Pre published/Staged/Develop/Multiple environment isolation
* [ ] Server: cluster and HA
* [ ] Security: ratelimiter
* [ ] Security: max http header/body limit
* [ ] Optimization: lightweight sidecar for distributed sync
* [ ] Optimization: file cache and small file read
* [ ] Web Management: web pages for management

## 3. Dependency

* [x] https://github.com/goproxy/goproxy - go module proxy
* [x] https://github.com/spf13/afero - file system support
* [x] https://github.com/BurntSushi/toml - toml config file parser
* [x] https://github.com/go-resty/resty - http client

Reference implementations

##### Go

* https://github.com/goproxy/goproxy
* https://github.com/goproxyio/goproxy

##### Maven

* https://github.com/dzikoysk/reposilite

## 4. Client support

#### 4.1 Go

* Sumdb: to avoid sumdb check for a private repository, use GONOSUMDB env. GOPRIVATE is an alternative if private proxy
  function is not enabled and used.

## A. Reference documents

### Go

[Go Module](https://go.dev/ref/mod)
[Go Module Sumdb Proxy](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#proxying-a-checksum-database)

### Cargo(Rust)

[Registries](https://doc.rust-lang.org/cargo/reference/registries.html)

### Nuget

[Nuget Server API Overview](https://learn.microsoft.com/en-us/nuget/api/overview)

### Docker

[Docker registry](https://docs.docker.com/registry/spec/api/)

### Maven

* [Maven Repository Centre](https://maven.apache.org/repositories/index.html)
* [Repository Management](https://maven.apache.org/repository-management.html)
* [Introduction to Repositories](https://maven.apache.org/guides/introduction/introduction-to-repositories.html)
* [Maven Metadata](https://maven.apache.org/repositories/metadata.html)

