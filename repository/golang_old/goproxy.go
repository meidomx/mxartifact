package golang_old

import "github.com/goproxy/goproxy"

func Init() {
	proxy := &goproxy.Goproxy{
		GoBinName:           "",
		GoBinEnv:            nil,
		GoBinMaxWorkers:     0,
		PathPrefix:          "",
		Cacher:              nil,
		CacherMaxCacheBytes: 0,
		ProxiedSUMDBs:       nil,
		Transport:           nil,
		TempDir:             "",
		ErrorLogger:         nil,
	}

	var _ = proxy

	//TODO
}
