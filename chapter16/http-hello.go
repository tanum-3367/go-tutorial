package main

import (
	"log"
	"net/http"
)

type hello struct {
}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", hello{}))
}
