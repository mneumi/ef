package init

import (
	"embed"
	"io/fs"
	"strings"

	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/core/fileop"
	"github.com/mneumi/ef/core/print"
	urface_cli "github.com/urfave/cli/v2"
)

//go:embed template/*
var templateFS embed.FS

func Command() *urface_cli.Command {
	return &urface_cli.Command{
		Name:  config.GetInit().Name,
		Usage: config.GetInit().Usage,
		Action: func(c *urface_cli.Context) error {
			exist, err := fileop.ExistPath(".ef")
			if err != nil {
				print.PrintlnError("判断 .ef 目录是否存在失败: %+v", err)
				return nil
			}
			if exist {
				print.PrintlnPrimary("当前目录已经是 ef 项目")
				return nil
			}

			err = initEf()
			if err != nil {
				print.PrintlnError("初始化 ef 项目失败: %+v", err)
			}

			print.PrintlnSuccess("ef 项目初始化成功")
			print.PrintlnPrimary("可以考虑将 .ef 目录加入到 .gitignore 文件中")

			return nil
		},
	}
}

func initEf() error {
	entities, err := fs.ReadDir(templateFS, "template")
	if err != nil {
		return err
	}

	readDirEntries(entities, "template")

	return nil
}

func readDirEntries(entries []fs.DirEntry, prevPath string) {
	for _, item := range entries {
		itemName := item.Name()
		next := prevPath + "/" + itemName

		if item.IsDir() {
			fileop.CreateDir(getRealPath(prevPath, itemName))

			newEntries, _ := fs.ReadDir(templateFS, next)

			readDirEntries(newEntries, next)
		} else {
			byte, _ := fs.ReadFile(templateFS, next)

			fileop.WriteFile(getRealPath(prevPath, itemName), byte)
		}
	}
}

func getRealPath(prevPath string, itemName string) string {
	result := prevPath + "/" + itemName
	result = strings.Replace(result, "template/", "", 1)

	return fileop.GetPathAbs(result)
}
