package cli

import (
	"os"

	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/command/create"
	"github.com/mneumi/ef/command/deploy"
	"github.com/mneumi/ef/command/gen"
	initial "github.com/mneumi/ef/command/init"
	"github.com/mneumi/ef/command/upgrade"
	"github.com/mneumi/ef/command/version"
	"github.com/mneumi/ef/core/print"

	urfave_cli "github.com/urfave/cli/v2"
)

type CLI struct {
	App *urfave_cli.App
}

func Start() {
	c := &CLI{
		App: &urfave_cli.App{},
	}

	c.bindInfo()
	c.bindCommand()

	c.Run()
}

func (c *CLI) Run() {
	err := c.App.Run(os.Args)
	if err != nil {
		print.PrintlnError("read config error: %v", err)
		os.Exit(1)
	}
}

func (c *CLI) bindInfo() {
	c.App.Name = config.GetInfo().Name
	c.App.Version = config.GetInfo().Version
	c.App.CustomAppHelpTemplate = config.GetInfo().Template
}

func (c *CLI) bindCommand() {
	commands := []*urfave_cli.Command{
		version.Command(),
		upgrade.Command(),
		create.Command(),
		deploy.Command(),
		initial.Command(),
		gen.Command(),
	}

	c.App.Commands = append(c.App.Commands, commands...)
}
