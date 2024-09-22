package docker

import "testing"

func TestUrlRegexp(t *testing.T) {
	if !end2Hash.Match([]byte("http://127.0.0.1/v2/myorg/myrepo/blobs/sha256:b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")) {
		t.Fatal("failed end2hash: http://127.0.0.1/v2/myorg/myrepo/blobs/sha256:b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9")
	}
	if !end2Hash.Match([]byte("http://127.0.0.1/v2/myorg/myrepo/blobs/sha256+b64u:LCa0a2j_xo_5m0U8HTBBNBNCLXBkg7-g-YpeiGJm564")) {
		t.Fatal("failed end2hash: http://127.0.0.1/v2/myorg/myrepo/blobs/sha256+b64u:LCa0a2j_xo_5m0U8HTBBNBNCLXBkg7-g-YpeiGJm564")
	}
	if end3Tag.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/.INVALID_MANIFEST_NAME") || end3Hash.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/.INVALID_MANIFEST_NAME") {
		t.Fatal("failed end3Tag & end3Hash: http://127.0.0.1/v2/myorg/myrepo/manifests/.INVALID_MANIFEST_NAME")
	}
	if !end3Hash.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/sha256:9dad4f699e3a49f674aeec428e0a973b8101800866569a3a9b96647c7a4fd238") || end3Tag.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/sha256:9dad4f699e3a49f674aeec428e0a973b8101800866569a3a9b96647c7a4fd238") {
		t.Fatal("failed: http://127.0.0.1/v2/myorg/myrepo/manifests/sha256:9dad4f699e3a49f674aeec428e0a973b8101800866569a3a9b96647c7a4fd238")
	}
	if !end3Hash.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/sha256+b64u:LCa0a2j_xo_5m0U8HTBBNBNCLXBkg7-g-YpeiGJm564") || end3Tag.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/sha256+b64u:LCa0a2j_xo_5m0U8HTBBNBNCLXBkg7-g-YpeiGJm564") {
		t.Fatal("failed: http://127.0.0.1/v2/myorg/myrepo/manifests/sha256+b64u:LCa0a2j_xo_5m0U8HTBBNBNCLXBkg7-g-YpeiGJm564")
	}
	if !end3Tag.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/tagtest0") || end3Hash.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/tagtest0") {
		t.Fatal("failed: http://127.0.0.1/v2/myorg/myrepo/manifests/tagtest0")
	}
	if !end3Tag.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/v1.3.2") || end3Hash.MatchString("http://127.0.0.1/v2/myorg/myrepo/manifests/v1.3.2") {
		t.Fatal("failed: http://127.0.0.1/v2/myorg/myrepo/manifests/v1.3.2")
	}
}
