package main

import (
	"log"
	"net/http"
	"time"

	"github.com/isaacd9/concurrency-middleware"
)

const addr = "localhost:8080"

func main() {
	handler := http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		log.Printf("%s %q 200", r.Method, r.URL)
		rw.Write([]byte("hello world"))
	})

	m := concurrency.NewMiddlware(concurrency.NewSyncSemaphore(10))

	log.Printf("starting server on %q", addr)
	if err := http.ListenAndServe(addr, m.Handle(handler)); err != nil {
		log.Fatalf("error: %+v", err)
	}
}
