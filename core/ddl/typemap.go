package ddl

var sqlTypeMap = map[string]string{
	"int":             "int",
	"tinyint":         "int",
	"bigint":          "int64",
	"bigint unsigned": "int64",
	"bool":            "bool",
	"varchar":         "string",
	"char":            "string",
	"text":            "string",
	"date":            "time.Time",
	"datetime":        "time.Time",
	"timestamp":       "time.Time",
	"time":            "time.Time",
	"decimal":         "float64",
}
