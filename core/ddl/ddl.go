package ddl

import (
	"errors"
	"regexp"

	"github.com/mneumi/ef/core/util"
	"github.com/xwb1989/sqlparser"
)

func SqlToMeta(sql string) (*Meta, error) {
	ddl, err := extractDDL(sql)

	if err != nil {
		return nil, err
	}

	return ddlToMeta(ddl), nil
}

func ddlToMeta(ddl *sqlparser.DDL) *Meta {
	tableName := extractTableName(ddl.NewName.Name.String())
	tableComment := extractTableComment(ddl.TableSpec.Options)

	primaryKeyColumns, indexColumns := extractIndexes(ddl.TableSpec.Indexes)

	fields := extractFields(ddl.TableSpec.Columns)

	return &Meta{
		TableName:         tableName,
		TableComment:      tableComment,
		PrimaryKeyColumns: primaryKeyColumns,
		IndexColumns:      indexColumns,
		Fields:            fields,
	}
}

func extractDDL(sql string) (*sqlparser.DDL, error) {
	sqlStmt, err := sqlparser.ParseStrictDDL(sql)
	if err != nil {
		return nil, err
	}

	ddlStmt, ok := sqlStmt.(*sqlparser.DDL)
	if !ok {
		return nil, errors.New("传入SQL无法解析为DDL，请检查SQL语句")
	}

	return ddlStmt, nil
}

func extractFields(columns []*sqlparser.ColumnDefinition) []*Field {
	var ret []*Field

	for _, item := range columns {
		goType, ok := sqlTypeMap[item.Type.Type]
		if !ok {
			goType = "未定义转换类型"
		}

		snakeName := item.Name.String()

		ret = append(ret, &Field{
			FieldName: &FieldName{
				SnakeName:  snakeName,
				PascalName: util.SnakeNameToPascal(snakeName),
				CamelName:  util.SnakeNameToCamelName(snakeName),
			},
			Type:    goType,
			NotNull: bool(item.Type.NotNull),
			Default: item.Type.Default.Val,
			Comment: string(item.Type.Comment.Val),
		})
	}

	return ret
}

func extractTableName(snakeName string) *TableName {
	return &TableName{
		SnakeName:  snakeName,
		PascalName: util.SnakeNameToPascal(snakeName),
		CamelName:  util.SnakeNameToCamelName(snakeName),
	}
}

func extractTableComment(input string) string {
	reg := regexp.MustCompile(`comment.+'(.+)'`)

	result := reg.FindStringSubmatch(input)

	if len(result) >= 2 {
		return result[1]
	}

	return ""
}

func extractIndexes(indexes []*sqlparser.IndexDefinition) ([]string, []string) {
	var primaryKeySlice []string
	var indexSlice []string

	for _, index := range indexes {
		switch index.Info.Type {
		case "primary key":
			for _, col := range index.Columns {
				primaryKeySlice = append(primaryKeySlice, col.Column.String())
			}
		case "index":
			for _, col := range index.Columns {
				indexSlice = append(indexSlice, col.Column.String())
			}
		}
	}

	return primaryKeySlice, indexSlice
}
