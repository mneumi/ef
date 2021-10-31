package config

type Create struct {
	Name  string     `yaml:"name"`
	Usage string     `yaml:"usage"`
	Group []string   `yaml:"group"`
	Type  []TypeItem `yaml:"type"`
}

type TypeItem struct {
	Name string `yaml:"name"`
	Desc string `yaml:"desc"`
}

func GetCreate() Create {
	return c.Create
}
