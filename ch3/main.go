package main

import (
	"database/sql"
	myfmt "fmt" //命名导入,给包指定别名,一般用于有相同名字的包的情况
	_ "github.com/go-sql-driver/mysql"
)

//一个查询mysql的例子
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		myfmt.Println(err)
	}
	rows, err := db.Query("select * from user where age > ?", 1)
	if err != nil {
		myfmt.Println(err)
	}
	for rows.Next() {
		var (
			id   int
			name string
			age  int
			sex  string
		)
		err := rows.Scan(&id, &name, &age, &sex)
		if err != nil {
		}
		myfmt.Printf("name:%s,age:%d,sex:%s \n", name, age, sex)
	}
}
