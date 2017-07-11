package main

import (
	"fmt"
	"./api"
)

func main() {
	api.Sqlite()
	fmt.Println("do sqlite ok")
	api.Mysql()
	fmt.Println("do mysql ok")

	userName, b := "saber", 2
	fmt.Println("userName=", userName)
	fmt.Println("b=", b)

	userName, c := "archer", 3
	fmt.Println("userName=", userName)
	fmt.Println("c=", c)
}