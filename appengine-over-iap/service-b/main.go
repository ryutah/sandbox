package main

import (
	"net/http"

	"log"

	"github.com/ryutah/sandbox/appengine-over-iap/lib"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Called!")
		w.Write([]byte("ServiceB!!"))
	})

	lib.Serve()
}
