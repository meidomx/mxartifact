package config

type Config struct {
	FileStorage struct {
		Location string `toml:"location"`
	} `toml:"file_storage"`

	Repository struct {
		Golang struct {
		} `toml:"golang"`
	} `toml:"repository"`
}
