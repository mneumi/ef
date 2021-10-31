package upgrade

import (
	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/core/cmd"
	"github.com/mneumi/ef/core/print"
	urface_cli "github.com/urfave/cli/v2"
)

func Command() *urface_cli.Command {
	return &urface_cli.Command{
		Name:  config.GetUpgrade().Name,
		Usage: config.GetUpgrade().Usage,
		Action: func(c *urface_cli.Context) error {
			_, err := cmd.RunCMD("go", "install", config.GetInfo().GitRepo)
			if err != nil {
				print.PrintlnError("升级失败: go install %+v", config.GetInfo().GitRepo)
				return nil
			}
			print.PrintlnSuccess("升级成功")
			return nil
		},
	}
}
