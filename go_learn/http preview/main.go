package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println(*r)

	w.Write([]byte("<h1>Hello from Go</h1>"))
}
