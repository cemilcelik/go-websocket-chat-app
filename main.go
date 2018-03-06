package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	fmt.Println("App started.")

	fs := http.FileServer(http.Dir("public"))

	http.Handle("/", fs)

	fmt.Println("Http server is running on http://localhost:8000")
	err := http.ListenAndServe("localhost:8000", fs)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
