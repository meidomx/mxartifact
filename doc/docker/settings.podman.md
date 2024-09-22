# Podman configuration

### Add mirrors

```text
nano /etc/containers/registries.conf


# by default this line is not required
unqualified-search-registries = ["docker.io"]

[[registry]]
prefix = "docker.io"
location = "mirror.gcr.io"
insecure = true

```
