package version

import (
	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/core/print"
	urface_cli "github.com/urfave/cli/v2"
)

func Command() *urface_cli.Command {
	return &urface_cli.Command{
		Name:  config.GetVersion().Name,
		Usage: config.GetVersion().Usage,
		Action: func(c *urface_cli.Context) error {
			print.PrintlnPrimary(config.GetInfo().Version)
			return nil
		},
	}
}
