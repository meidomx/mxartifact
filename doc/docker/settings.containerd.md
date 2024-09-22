# Containerd configuration

### Add mirrors for v1.x

```text
nano /etc/containerd/config.toml

    [plugins."io.containerd.grpc.v1.cri".registry]
      config_path = "/etc/containerd/certs.d"

```

```text
nano /etc/containerd/certs.d/docker.io/hosts.toml

server = "https://mirror.gcr.io"
skip_verify = true
[host."https://mirror.gcr.io"]
  capabilities = ["pull", "resolve"]
  skip_verify = true

```
