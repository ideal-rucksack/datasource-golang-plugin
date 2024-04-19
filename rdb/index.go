package rdb

type Client interface {
	// Databases 获取数据库列表
	Databases() ([]string, error)

	// Tables 获取表列表
	Tables(database string) ([]string, error)

	// Columns 获取列列表
	Columns(database, table string) ([]string, error)

	// ExecuteQuery 查询
	ExecuteQuery(databaseName string, query string) ([]map[string]interface{}, error)

	// Execute 执行
	Execute(database, sql string) error

	// Ping 测试连接
	Ping() error

	// TableSchema 获取表结构
	TableSchema(database, table string) (TableSchema, error)
}

// TableSchema defines the schema of a database table
type TableSchema struct {
	Columns []ColumnSchema
}

// ColumnSchema defines the schema of a column in a database table
type ColumnSchema struct {
	Name         string
	DataType     string
	IsNullable   bool
	IsPrimaryKey bool
}
