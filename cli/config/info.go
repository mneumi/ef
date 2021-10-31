package config

type Info struct {
	Name     string `yaml:"name"`
	Version  string `yaml:"version"`
	GitRepo  string `yaml:"gitrepo"`
	Template string `yaml:"template"`
}

func GetInfo() *Info {
	return &c.Info
}
