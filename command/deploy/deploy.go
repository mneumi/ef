package deploy

import (
	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/core/cmd"
	"github.com/mneumi/ef/core/print"
	urface_cli "github.com/urfave/cli/v2"
)

func Command() *urface_cli.Command {
	return &urface_cli.Command{
		Name:  config.GetDeploy().Name,
		Usage: config.GetDeploy().Usage,
		Action: func(c *urface_cli.Context) error {
			_, err := cmd.RunCMD("stark-tools", "deploy")
			if err != nil {
				print.PrintlnError("部署失败: %+v", err)
				return nil
			}
			print.PrintlnSuccess("开始部署")
			return nil
		},
	}
}
