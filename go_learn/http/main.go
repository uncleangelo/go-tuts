package main

import (
	"fmt"
	"net/http"
)

type handsome int

func (h handsome) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request--------------------------\n", r)
	w.Header().Set("Content-Type", "text/html")
	// w.Write([]byte("<h1>Hello</h1>"))
	// io.Copy(w, strings.NewReader("<h4>Heyheyhey</h4>"))
	fmt.Fprint(w, "<a href=\"http://google.com\">google</a>")
}

func main() {
	var h handsome
	http.ListenAndServe(":9000", h)
}
