package example

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// GetDSN retrieves dsn from viper config.
func GetDSN() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&multiStatements=true&interpolateParams=true",
		"root", "dc", "127.0.0.1", 3306, "daycam",
	)
}

// 如何从 sql 抽象中间层 -> driver.Query 层
// sql: http://luodw.cc/2016/08/28/golang02/
func TestExample1(t *testing.T) {
	// GetDSN 方法，看看
	// DSN = data source name
	// 1. 根据 driverName, 获取提前注册的 driver
	// 2. 初始化一个一个数据库实例 db
	// 2.1 如果 driverContext 有实现, 则通过 driver.OpenConnector 初始化
	// 2.2 否则通过 DSNConnector 进行初始化
	// 最后通过调用 OpenDB 完成初始化工作
	db, err := sql.Open("mysql", GetDSN())
	checkErr(err)

	// 1. 如何检查参数输入，是否符合条件的？  SELECT * FROM user WHERE user_id = ? ( 但是没有传 user_id 的情况？)
	rows, err := db.Query("select * from test")
	checkErr(err)
	var (
		id   uint64
		name string
	)
	for rows.Next() {
		err = rows.Scan(&id, &name)
		checkErr(err)
		fmt.Println(id, "->", name)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
