package config

type Init struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`
}

func GetInit() Init {
	return c.Init
}
