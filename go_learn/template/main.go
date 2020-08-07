package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":9002", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	template, _ := template.ParseFiles("index.htm")
	data := struct {
		ID   int
		Name string
	}{
		1,
		"John",
	}
	template.Execute(w, data)
	// fmt.Fprint(w, "asdsafasf")
}
