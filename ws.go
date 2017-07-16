package main

import (
	"golang.org/x/net/websocket"
	"fmt"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("recieve ws data: " + reply)

		msg := "Response data:" + reply + "order list"
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
			break
		}
	}
}

func main() {
	fmt.Println("websocket backend begin")
	http.Handle("/", http.FileServer(http.Dir("."))) // <-- note this line

	http.Handle("/websocket", websocket.Handler(Echo))

	if err := http.ListenAndServe(":8083", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	fmt.Println("websocket backend end")
}