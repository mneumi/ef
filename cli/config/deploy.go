package config

type Deploy struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`
}

func GetDeploy() Deploy {
	return c.Deploy
}
