package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "world")
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}
