package config

type Gen struct {
	Name    string `yaml:"name"`
	Usage   string `yaml:"usage"`
	Modules struct {
		Pb    Pb    `yaml:"pb"`
		Model Model `yaml:"model"`
	}
}

type Pb struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`
}

type Model struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`
}

func GetGen() Gen {
	return c.Gen
}
