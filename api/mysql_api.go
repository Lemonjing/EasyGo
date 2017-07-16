package api

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Mysql() {
	insert()
}

// insert
func insert() {
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

func query() {
	db, err := sql.Open("mysql", "root:root@/golang")
	checkErr(err)
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	for rows.Next() {
		var userId int
		var userName string
		var userAge int
		var userSex int
		rows.Columns()
		err = rows.Scan(&userId, &userName, &userAge, &userSex)
		checkErr(err)
		fmt.Println("userId=", userId)
		fmt.Println("userName=", userName)
		fmt.Println("userAge=", userAge)
		fmt.Println("userSex=", userSex)
	}
}

//更新数据
func update() {
	db, err := sql.Open("mysql", "root:root@/golang")
	checkErr(err)

	stmt, err := db.Prepare(`UPDATE user SET user_name=?,user_sex=? WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec("saber", 2, 1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

//删除数据
func remove() {
	db, err := sql.Open("mysql", "root:root@/golang")
	checkErr(err)

	stmt, err := db.Prepare(`DELETE FROM user WHERE user_id=?`)
	checkErr(err)
	res, err := stmt.Exec(1)
	checkErr(err)
	num, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(num)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
