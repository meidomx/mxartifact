# Docker configuration

### Add mirrors

```text
cat /etc/docker/daemon.json

{
  "registry-mirrors": ["https://mirror.gcr.io"],
  "insecure-registries": ["mirror.gcr.io"]
}
```

