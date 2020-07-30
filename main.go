package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var addr = ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	log.Println("listening on", addr)
	http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if sleep := r.URL.Query().Get("sleep"); sleep != "" {
			d, err := time.ParseDuration(sleep)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			time.Sleep(d)
		}
		bts, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		log.Println(r.URL, "body:", string(bts))
		fmt.Fprintln(w, "ok")
	}))
}
