package config

type Version struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`
}

func GetVersion() Version {
	return c.Version
}
