package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// reader writer
	// reader - pembaca X yang boleh dibaca  (readable)
	// writer - penulis X yang boleh ditulis (writable)

	file, err := os.Open("a1")
	if err != nil {
		log.Fatalln(err)
	}

	// text, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	io.Copy(os.Stdout, file)
}
