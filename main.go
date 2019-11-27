package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var addr = ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	log.Println("listening on", addr)
	http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
		log.Println(r.URL)
	}))
}
