package config

type Config struct {
	FileStorage struct {
		Location string `toml:"location"`
	} `toml:"file_storage"`

	Repository struct {
		Golangs []GoRepoConf `toml:"golang"`
	} `toml:"repository"`

	Shared struct {
		Debug  bool   `toml:"debug"`
		Listen string `toml:"listen"`
	} `toml:"mxartifact"`
}

type GoRepoConf struct {
	Name               string `toml:"name"`
	Type               string `toml:"type"`
	BaseUrl            string `toml:"base_url"`
	HttpProxy          string `toml:"http_proxy"`
	UpstreamRepository string `toml:"upstream_repository"`

	ParentRepository string `toml:"parent_repository"`
}
