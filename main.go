package main

import (
	"fmt"
	"log"
	"net/http"
)

const addr = ":8111"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request from", r.RemoteAddr)
		fmt.Fprintf(w, "hello, world")
	})
	log.Println("listening on", addr)
	http.ListenAndServe(addr, nil)
}
