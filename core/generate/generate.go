package generate

import (
	"strings"
	"text/template"

	"github.com/mneumi/ef/core/cmd"
	"github.com/mneumi/ef/core/ddl"
	"github.com/mneumi/ef/core/fileop"
	"github.com/mneumi/ef/templates/mvc"
)

func GenModel(path string, meta *ddl.Meta) error {
	tmpl, err := template.New("model").Parse(mvc.ModelTpl)
	if err != nil {
		return err
	}

	relativePath := strings.Join([]string{"model", path}, fileop.GetSeparator())

	absPath := fileop.GetPathAbs(relativePath)

	err = fileop.CreateDirRelative(relativePath)
	if err != nil {
		return err
	}

	f, err := fileop.OpenFile(absPath, fileop.O_CREATE|fileop.O_TRUNC)
	if err != nil {
		return err
	}

	err = tmpl.Execute(f, meta)
	if err != nil {
		return err
	}

	err = formatGoFile(absPath)
	if err != nil {
		return err
	}

	return nil
}

func formatGoFile(path string) error {
	_, err := cmd.RunCMD("go", "fmt", path)
	return err
}
