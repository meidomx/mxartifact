# Podman configuration

### Add mirrors

```text
nano /etc/containers/registries.conf


unqualified-search-registries = ["docker.io"]

[[registry]]
prefix = "docker.io"
location = "mirror.gcr.io"
insecure = true

```
