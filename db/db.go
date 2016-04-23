package db

func Connect(drivername, username, password, hostname string, port int, databasename string) error {
	conn, err := dbGetConnect(drivername, username, password, hostname, port, databasename)
	if err != nil {
		return err
	} else {
		db = conn
		return nil
	}
}

//值设置器
type ISet interface {
	Values(values ...interface{}) (int64, error)
}

//单行读取
type IRow interface {
	Scan(dest ...interface{}) error
	Struct(object interface{}) error

	Slice() ([]interface{}, error)
	Map() (map[string]interface{}, error)
}

//多行读取
type IRows interface {
	IRow
	Close() error
	Err() error
	Next() bool
	Columns() ([]string, error)
}

//表操作接口
type ITable interface {

	//输出表结构
	ToSql() string

	//添加记录
	Add(values ...interface{}) (int64, error)

	//删除记录
	Del(args ...interface{}) (int64, error)

	//更新记录
	Update(args ...interface{}) ISet

	//获取单条记录
	Get(args ...interface{}) IRow

	//查询多条记录
	Find(args ...interface{}) (IRows, error)
	List(take, skip int) (IRows, error)
	Query(query string, args ...interface{}) (IRows, error)

	//统计
	Count(query string, args ...interface{}) (int64, error)
}
