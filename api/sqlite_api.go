package main

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	//创建表
	sql_table_create := `
    CREATE TABLE IF NOT EXISTS user(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
        created DATE NULL
    );
    CREATE TABLE IF NOT EXISTS topic(
    	tid INTEGER PRIMARY KEY AUTOINCREMENT,
        topicname VARCHAR(64) NULL,
        username VARCHAR(64) NULL,
        created DATE NULL
    )
    `
	db.Exec(sql_table_create)

	// insert
	stmt1, err := db.Prepare("INSERT INTO user(username, departname, created) values(?,?,?)")
	stmt2, err := db.Prepare("INSERT INTO topic(topicname, username, created) values(?,?,?)")
	checkErr(err)

	res1, err := stmt1.Exec("陶然", "滴滴", "2017-07-01")
	res2, err := stmt2.Exec("topic", "saber", "2017-07-01")
	checkErr(err)

	id1, err := res1.LastInsertId()
	id2, err := res2.LastInsertId()
	checkErr(err)

	fmt.Println("id1=", id1)
	fmt.Println("id2=", id2)

	// update
	stmt1, err = db.Prepare("update user set username=? where uid=?")
	checkErr(err)

	res1, err = stmt1.Exec("陶然_new", id1)
	checkErr(err)

	affect, err := res1.RowsAffected()
	checkErr(err)

	fmt.Println("affect=", affect)

	// query
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	var uid int
	var username string
	var department string
	var created time.Time

	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("query_start")
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
		fmt.Println("query_end")
	}

	rows.Close()

	// query2
	rows2, err := db.Query("SELECT * FROM topic")
	checkErr(err)
	var tid int
	var topicName string
	var userName string
	var createdTime time.Time

	for rows2.Next() {
		err = rows2.Scan(&tid, &topicName, &userName, &createdTime)
		checkErr(err)
		fmt.Println("query_start2")
		fmt.Println(tid)
		fmt.Println(topicName)
		fmt.Println(userName)
		fmt.Println(createdTime)
		fmt.Println("query_end2")
	}

	rows2.Close()

	// delete
	stmt1, err = db.Prepare("delete from user where uid=?")
	checkErr(err)

	res1, err = stmt1.Exec(id1)
	checkErr(err)

	affect, err = res1.RowsAffected()
	checkErr(err)

	fmt.Println("affect=", affect)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}