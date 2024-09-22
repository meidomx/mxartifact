# General mirror settings

### DNS configure

**nekoq-bootstrap** configure example:

```text
"docker.registry.internal"="10.11.0.7"
"gcr.registry.internal"="10.11.0.7"
"k8sgcr.registry.internal"="10.11.0.7"
"ghcr.registry.internal"="10.11.0.7"
"quay.registry.internal"="10.11.0.7"
"k8sregistry.registry.internal"="10.11.0.7"

```

### Containerd configure

```text
mkdir -p /etc/containerd/certs.d/docker.io
cat <<EOF > /etc/containerd/certs.d/docker.io/hosts.toml
server = "http://docker.registry.internal"
skip_verify = true
[host."http://docker.registry.internal"]
  capabilities = ["pull", "resolve"]
  skip_verify = true
EOF

mkdir -p /etc/containerd/certs.d/gcr.io
cat <<EOF > /etc/containerd/certs.d/gcr.io/hosts.toml
server = "http://gcr.registry.internal"
skip_verify = true
[host."http://gcr.registry.internal"]
  capabilities = ["pull", "resolve"]
  skip_verify = true
EOF

mkdir -p /etc/containerd/certs.d/k8s.gcr.io
cat <<EOF > /etc/containerd/certs.d/k8s.gcr.io/hosts.toml
server = "http://k8sgcr.registry.internal"
skip_verify = true
[host."http://k8sgcr.registry.internal"]
  capabilities = ["pull", "resolve"]
  skip_verify = true
EOF

mkdir -p /etc/containerd/certs.d/ghcr.io
cat <<EOF > /etc/containerd/certs.d/ghcr.io/hosts.toml
server = "http://ghcr.registry.internal"
skip_verify = true
[host."http://ghcr.registry.internal"]
  capabilities = ["pull", "resolve"]
  skip_verify = true
EOF

mkdir -p /etc/containerd/certs.d/quay.io
cat <<EOF > /etc/containerd/certs.d/quay.io/hosts.toml
server = "http://quay.registry.internal"
skip_verify = true
[host."http://quay.registry.internal"]
  capabilities = ["pull", "resolve"]
  skip_verify = true
EOF

mkdir -p /etc/containerd/certs.d/registry.k8s.io
cat <<EOF > /etc/containerd/certs.d/registry.k8s.io/hosts.toml
server = "http://k8sregistry.registry.internal"
skip_verify = true
[host."http://k8sregistry.registry.internal"]
  capabilities = ["pull", "resolve"]
  skip_verify = true
EOF

```

### Podman configure

```text
unqualified-search-registries = ["docker.io"]

[[registry]]
prefix = "docker.io"
location = "docker.registry.internal"
insecure = true

[[registry]]
prefix = "gcr.io"
location = "gcr.registry.internal"
insecure = true

[[registry]]
prefix = "k8s.gcr.io"
location = "k8sgcr.registry.internal"
insecure = true

[[registry]]
prefix = "ghcr.io"
location = "ghcr.registry.internal"
insecure = true

[[registry]]
prefix = "quay.io"
location = "quay.registry.internal"
insecure = true

[[registry]]
prefix = "registry.k8s.io"
location = "k8sregistry.registry.internal"
insecure = true

```
