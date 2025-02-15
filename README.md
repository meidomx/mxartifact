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
    * [ ] Repository - Maven: client support
        * Maven
        * Gradle
* [x] Repository: Docker
    * [x] OCI Distribution Spec version: v1.0.1
    * [x] Repository - Docker: docker pull image support
    * [x] Repository - Docker: support binding default http port and dispatching via hostname
    * [ ] Repository - Docker: support binding default https port and dispatching via hostname
    * [x] Repository - Docker: upstream type: proxy server
        * Support pull from
            * docker.io / registry-1.docker.io
            * gcr.io
            * k8s.gcr.io
            * ghcr.io
            * quay.io
            * registry.k8s.io
        * Via mirror configuration in container client
    * [x] Repository - Docker: verified support for 3 ways of docker pull:
        * server/org/repo , server:port/org/repo
        * org/repo => docker.io/org/repo
        * repo => docker.io/library/repo
        * default tag without specifying one - latest
    * [x] Repository - Docker: client support
        * docker
        * podman
        * containerd
* [x] Upstream: http proxy support
* [x] Customizable Repository: layered(nested and recursive) repositories and file permission
* [x] Customizable Repository: separate url for each repository

### 1.1 Golang

#### Prerequisite

* Install `golang >= 1.23`

## 2. TODO List

* [ ] Repository: Nuget
* [ ] Repository: Cargo
* [ ] Upstream: Queuing
* [ ] Repository - Maven: auth and permissions
* [ ] Repository - Maven: repository type - release / snapshot / etc.
* [ ] Repository - Maven: local storage
* [ ] Repository - Maven: deploy with auth - username & password / other methods
* [ ] Repository - Maven: metadata for deploying snapshots
* [ ] Repository - Maven: Http Headers for GET/HEAD/PUT/POST/DELETE requests
* [ ] Repository - Go: need extract request type before processing repositories & rely on the type
* [ ] Repository - Go: support safe concurrent local storage
* [ ] Repository - Go: simulating GOPRIVATE
* [ ] Repository - Go: upstream - A.source repository C.internal recursive repository
* [ ] Repository - Go: checksum database
* [ ] Repository - Go: Max file size defined in the Go Module specification
* [ ] Repository - Go: support management for local persisted repositories including uploading/deleting files
* [ ] Repository - Go: access token for pulling private repository
* [ ] Repository - Docker: Support user and permission management, aka auth
    * [ ] Support authentication header transmission between client and registry & auto auth token handling.
        * hub.docker.com relying on the mechanism
* [ ] Repository - Docker: Support local persistent
* [ ] Repository - Docker: special handling of tag: latest
* [ ] Repository - Docker: OCI Distribution Spec version: v1.1.0 & Docker image manifest version 2, schema 2
* [ ] Repository - Docker: replicate image to another registry
* [ ] Repository: streaming transfer for large files in slow upstream env
* [ ] Repository: helm, apt
* [ ] Workflow support: Production/Pre published/Staged/Develop/Multiple environment isolation
* [ ] Server: cluster and HA
* [ ] Server: resource manager and reuse support including http listening addresses for multiplexing of docker and other
  repository types
* [ ] Server: make sure http base urls mutual exclusion
* [ ] Security: ratelimiter
* [ ] Security: max http header/body limit
* [ ] Optimization: Support streaming file download - reduce resource cost, enable download progressbar support,
  multi-layer simultaneously downloading
* [ ] Optimization: lightweight sidecar for distributed sync
* [ ] Optimization: file cache and small file read
* [ ] Web Management: web pages for management
* [ ] File Storage: S3
* [ ] File Storage: Multiple file storage
* [ ] Debugging: more logs for troubleshooting

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

* [Docker registry](https://docs.docker.com/registry/spec/api/)
* [Open Container Initiative Distribution Specification v1.0.1](https://github.com/opencontainers/distribution-spec/blob/v1.0.1/spec.md)
* [Open Container Initiative Distribution Specification v1.1.0](https://github.com/opencontainers/distribution-spec/blob/v1.1.0/spec.md)
* [Distribution Spec Protocol Detail v1.0.1](https://github.com/opencontainers/distribution-spec/blob/v1.0.1/detail.md)
* Registry test tool: https://github.com/opencontainers/distribution-spec/blob/v1.0.1/conformance
* [Image Manifest V 2, Schema 2](https://distribution.github.io/distribution/spec/manifest-v2-2/)

##### Docker registry special handling

**Pulling image behavior verification**

* Original docker client while pulling image from hub.docker.com:
    * Step1: request: docker.io - (need to check purpose)
    * Step2: request: registry-1.docker.io
    * Step3: request: auth.docker.io
* Docker + setup mirror site:
    * example mirror site: https://mirror.gcr.io provided
      by https://cloud.google.com/artifact-registry/docs/pull-cached-dockerhub-images
    * Step1: request: docker.io - (need to check purpose)
    * Step2: request: mirror site - mirror.gcr.io
* ctr from containerd + pulling image from hub.docker.com:
    * Step1: request: registry-1.docker.io
    * Step2: request: auth.docker.io
* ctr from containerd + setup mirror site:
    * example mirror site: https://mirror.gcr.io and
      configuration https://github.com/containerd/containerd/blob/main/docs/cri/registry.md
    * Doc ref2: https://github.com/containerd/containerd/blob/main/docs/cri/config.md#registry-configuration
    * Doc ref3: https://github.com/containerd/containerd/blob/main/docs/hosts.md
    * Step1: request: mirror site - mirror.gcr.io

Note:

* For the request of auth.docker.io, it is triggered by the response where the header `Www-Authenticate` exists.
    * The request details is based on the header which includes service, scope

**Requirement of authentication while pulling images from hub.docker.com**

* Support http header from registry. This header will indicate where and what to do auth

```text
Www-Authenticate:[Bearer realm="https://auth.docker.io/token",service="registry.docker.io",scope="repository:library/nginx:pull"]
```

* Support auth headers transferring between client and registry

* Auth ref1: https://distribution.github.io/distribution/spec/auth/jwt/
* Auth ref2: https://docs.docker.com/docker-hub/download-rate-limit/

**Ways to mirror hub.docker.com via default user experience(e.g. docker pull <image_name>)**

* Use mirror for hub.docker.com
    * Just standard requests to retrieve images
    * Mirror services are required to follow the behavior of hub.docker.com to do auth
    * Support both secure and insecure access to mirror
* Pure http proxy to docker client
    * Transparent access to hub.docker.com
    * Cannot support image caching since the traffic is encrypted by default
* Proxy docker.io related domains to mirror server with both registry APIs and auth APIs
    * Support caching images
    * Need support TLS on mirror server
    * May need to insecure the docker.io registry since the mirror server may not have valid cert
    * Also need to setup dns to respond mirror server addresses

### Maven

* [Maven Repository Centre](https://maven.apache.org/repositories/index.html)
* [Repository Management](https://maven.apache.org/repository-management.html)
* [Introduction to Repositories](https://maven.apache.org/guides/introduction/introduction-to-repositories.html)
* [Maven Metadata](https://maven.apache.org/repositories/metadata.html)

