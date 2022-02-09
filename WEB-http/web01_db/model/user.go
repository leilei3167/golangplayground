package model

import (
	"GOproject/src/go_code/WEB-http/web01_db/utils"
	"fmt"
)

//实现对数据库的增删改查
type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

//添加User的方法一:预编译后执行,可防止sql注入
func (user *User) AddUser() error {
	//向数据库中插入

	//1.sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	//2.预编译,Prepare创建一个准备好的*sql.Stmt状态用于之后的查询和命令
	//返回值可以同时执行多个查询和命令。
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译异常:", err)
	}
	//*Stmt就可以调用方法执行操作

	//3.执行
	_, err = inStmt.Exec("admin", "123456", "1356556043@qq.com")
	if err != nil {
		fmt.Println("方法一执行异常:", err)
	}
	return nil
}

//添加User的方法二:不用prepare
func (user *User) AddUser2() error {
	//向数据库中插入

	//1.sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	//2.执行,直接用DB的方法执行
	_, err := utils.Db.Exec(sqlStr, "admin2", "666666", "123@qq.com")
	if err != nil {
		fmt.Println("方法二执行异常:", err)
	}
	return nil
}
