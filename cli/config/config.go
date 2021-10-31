package config

import (
	_ "embed"
	"os"

	"github.com/mneumi/ef/core/print"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Info    Info    `yaml:"info"`
	Version Version `yaml:"version"`
	Upgrade Upgrade `yaml:"upgrade"`
	Deploy  Deploy  `yaml:"deploy"`
	Init    Init    `yaml:"init"`
	Create  Create  `yaml:"create"`
	Gen     Gen     `yaml:"gen"`
}

var c Config

//go:embed config.yaml
var configByte []byte

func init() {
	err := yaml.Unmarshal(configByte, &c)
	if err != nil {
		print.PrintlnError("read config error: %v", err)
		os.Exit(1)
	}
}
