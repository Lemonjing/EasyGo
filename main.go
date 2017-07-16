package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

func main() {
	fmt.Println("==========")
	db, err := sql.Open("mysql", "root:root@/golang")
	checkErr(err)

	stmt, err := db.Prepare(`INSERT user (user_name,user_age,user_sex) values (?,?,?)`)
	checkErr(err)
	res, err := stmt.Exec("tony", 20, 1)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("id=", id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}