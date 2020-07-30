package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var addr = ":8080"
	if p := os.Getenv("PORT"); p != "" {
		addr = ":" + p
	}
	log.Println("listening on", addr)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
	})

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Println("server started")

	<-done
	log.Println("server stop requested")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed:%+v", err)
	}
	log.Print("server exit properly")
}
