package pb

import (
	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/core/cmd"
	"github.com/mneumi/ef/core/print"
	urface_cli "github.com/urfave/cli/v2"
)

func Command() *urface_cli.Command {
	return &urface_cli.Command{
		Name:  config.GetGen().Modules.Pb.Name,
		Usage: config.GetGen().Modules.Pb.Usage,
		Action: func(c *urface_cli.Context) error {
			_, err := cmd.RunCMD("stark-tools", "genpb")
			if err != nil {
				print.PrintlnError("pb生成失败: %+v", err)
				return nil
			}
			print.PrintlnSuccess("pb生成成功")
			return nil
		},
	}
}
