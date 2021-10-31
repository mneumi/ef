package model

import (
	"io/fs"
	"strings"

	"github.com/mneumi/ef/cli/config"
	"github.com/mneumi/ef/core/ddl"
	"github.com/mneumi/ef/core/fileop"
	"github.com/mneumi/ef/core/generate"
	"github.com/mneumi/ef/core/print"

	urface_cli "github.com/urfave/cli/v2"
)

func Command() *urface_cli.Command {
	return &urface_cli.Command{
		Name:  config.GetGen().Modules.Model.Name,
		Usage: config.GetGen().Modules.Model.Usage,
		Action: func(c *urface_cli.Context) error {
			root := getRoot()

			exist, _ := fileop.ExistPath(root)
			if !exist {
				print.PrintlnError("检查当前目录下是否存在目录 %+v", root)
				return nil
			}

			m, err := extractSqlInfo(root)
			if err != nil {
				print.PrintlnError("读取SQL信息失败 %+v", err)
				return nil
			}

			err = genModleFile(m)
			if err != nil {
				print.PrintlnError("生成 Model 文件失败 %+v", err)
				return nil
			}

			return nil
		},
	}
}

func getRoot() string {
	return strings.Join([]string{".ef", "model"}, fileop.GetSeparator())
}

func extractSqlInfo(root string) (map[string]string, error) {
	m := make(map[string]string)

	modelDir := fileop.GetPathAbs(root)

	fileop.Walk(modelDir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			ext := fileop.GetFileExt(info.Name())
			if ext == "sql" {
				sqlContent, err := fileop.ReadFile(path)
				if err != nil {
					return err
				}
				shortPath := strings.Split(path, root)[1]
				m[shortPath] = sqlContent
			}
		}

		return nil
	})

	return m, nil
}

func genModleFile(m map[string]string) error {
	for path, sqlContent := range m {
		meta, err := ddl.SqlToMeta(sqlContent)
		if err != nil {
			return err
		}

		goFilePath := strings.ReplaceAll(path, ".sql", ".go")

		err = generate.GenModel(goFilePath, meta)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
