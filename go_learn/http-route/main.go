package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>Home</h1><br><a href=\"/web\">Web</a>")
	})
	mux.HandleFunc("/web", webpage)

	http.ListenAndServe(":9000", mux)
}

func webpage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<strong>Strong</strong>")
}
