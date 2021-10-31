package gen

import (
	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/command/gen/model"
	"github.com/mneumi/ef/command/gen/pb"
	"github.com/mneumi/ef/core/print"
	urface_cli "github.com/urfave/cli/v2"
)

func Command() *urface_cli.Command {
	return &urface_cli.Command{
		Name:  config.GetGen().Name,
		Usage: config.GetGen().Usage,
		Action: func(c *urface_cli.Context) error {
			print.PrintlnPrimary("TODO 生成全部")
			return nil
		},
		Subcommands: []*urface_cli.Command{
			pb.Command(),
			model.Command(),
		},
	}
}
