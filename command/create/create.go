package create

import (
	"fmt"
	"os"

	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/core/cmd"
	"github.com/mneumi/ef/core/print"
	"github.com/mneumi/ef/core/prompt"

	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  config.GetCreate().Name,
		Usage: config.GetCreate().Usage,
		Action: func(c *cli.Context) error {
			projectName := getProjectName()
			projectType := getProjectType()
			projectGroup := getProjectGroup()

			_, err := cmd.RunCMD("stark-tools",
				"new", projectType,
				"-g", projectGroup,
				"-n", projectName)
			if err != nil {
				print.PrintlnError("新项目创建失败: %+v", err)
				return nil
			}

			print.PrintlnSuccess("新项目 %+v 创建成功！", projectName)

			return nil
		},
	}
}

func getProjectName() string {
	projectName, _ := prompt.Text("请输入项目名称", func(input string) error {
		return nil
	})

	if projectName == "" {
		print.PrintlnError("项目名称不能为空")
		os.Exit(1)
	}

	return projectName
}

func getProjectType() string {
	var groups []prompt.PromptItem

	for _, item := range config.GetCreate().Type {
		groups = append(groups, prompt.PromptItemPrimary(fmt.Sprintf("%s : %s", item.Name, item.Desc)))
	}

	groups = append(groups, "取消")

	idx, _, err := prompt.Select("请选择新建的项目类型", groups)

	if err != nil {
		print.PrintlnError("选择新建的项目类型失败：%+v", err)
		os.Exit(1)
	}

	if idx == len(groups)-1 {
		os.Exit(1)
	}

	return config.GetCreate().Type[idx].Name
}

func getProjectGroup() string {
	var groups []prompt.PromptItem

	for _, item := range config.GetCreate().Group {
		groups = append(groups, prompt.PromptItemPrimary(item))
	}

	groups = append(groups, prompt.PromptItemPrimary("取消"))

	idx, result, err := prompt.Select("请选择对应的项目分组", groups)

	if err != nil {
		print.PrintlnError("选择对应的项目分组失败：%+v", err)
		os.Exit(1)
	}

	if idx == len(groups)-1 {
		os.Exit(1)
	}

	return result
}
