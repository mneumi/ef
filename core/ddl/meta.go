package ddl

type Meta struct {
	TableName         *TableName
	TableComment      string
	PrimaryKeyColumns []string
	IndexColumns      []string
	Fields            []*Field
}

type TableName struct {
	SnakeName  string
	CamelName  string
	PascalName string
}

type Field struct {
	FieldName *FieldName
	Type      string
	NotNull   bool
	Default   interface{}
	Comment   string
}

type FieldName struct {
	SnakeName  string
	CamelName  string
	PascalName string
}
