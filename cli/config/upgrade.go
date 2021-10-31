package config

type Upgrade struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`
}

func GetUpgrade() Upgrade {
	return c.Upgrade
}
