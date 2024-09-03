package config

type Config struct {
	FileStorage struct {
		Location string `toml:"location"`
	} `toml:"file_storage"`

	Repository struct {
		Golangs []GoRepoConf     `toml:"golang"`
		Mavens  []MavenRepoConf  `toml:"maven"`
		Dockers []DockerRepoConf `toml:"docker"`
	} `toml:"repository"`

	Shared struct {
		Debug     bool               `toml:"debug"`
		Listen    string             `toml:"listen"`
		Listeners []ListenerElemConf `toml:"listeners"`
	} `toml:"mxartifact"`
}

type ListenerElemConf struct {
	Name      string            `toml:"name"`
	Addresses []string          `toml:"addresses"`
	Options   map[string]string `toml:"options"`
}

type GoRepoConf struct {
	Name               string `toml:"name"`
	Type               string `toml:"type"`
	BaseUrl            string `toml:"base_url"`
	HttpProxy          string `toml:"http_proxy"`
	UpstreamRepository string `toml:"upstream_repository"`

	ParentRepository string `toml:"parent_repository"`
}

type MavenRepoConf struct {
	Name               string `toml:"name"`
	Type               string `toml:"type"`
	BaseUrl            string `toml:"base_url"`
	HttpProxy          string `toml:"http_proxy"`
	UpstreamRepository string `toml:"upstream_repository"`

	ParentRepository string `toml:"parent_repository"`
}

type DockerRepoConf struct {
	Name               string `toml:"name"`
	Type               string `toml:"type"`
	BaseUrl            string `toml:"base_url"`
	HttpProxy          string `toml:"http_proxy"`
	UpstreamRepository string `toml:"upstream_repository"`
	BindListeners      []struct {
		Name    string            `toml:"name"`
		Options map[string]string `toml:"options"`
	} `toml:"bind_listeners"`

	ParentRepository string `toml:"parent_repository"`
}
