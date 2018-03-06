package main

import (
	"fmt"
	"net/http"
)

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
