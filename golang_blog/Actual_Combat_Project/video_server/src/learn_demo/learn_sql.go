package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
    db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/video_db?charset=utf8")
    if err != nil {
    	panic(err)
	}
	//插入数据
    //stmt, err := db.Prepare("INSERT users SET login_name =?, pwd=?")
    //res, err := stmt.Exec("zheng", "aaaaaa")
    //id, err := res.LastInsertId()
    //if err != nil {
    //	panic(err)
	//}
	//fmt.Println(id)
	//更新操作
	//stmt, err := db.Prepare("UPDATE users set login_name=? where id=?")
	//res, err := stmt.Exec("linsijie", 2)
	//id, err := res.RowsAffected()
	//if err != nil {
	//	panic(err)
	//}
    //fmt.Println(id)
    //查询
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
    	panic(err)
	}
    for rows.Next() {
    	var id int
    	var login_name string
    	var pwd string
    	err = rows.Scan(&id, &login_name, &pwd)
    	fmt.Println(id, login_name, pwd)
	}

	//删除
	delete_stmt, err := db.Prepare("delete from users where id=?")
	res, err = delete_stmt.Exec(3)

}