package lib

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Serve() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
