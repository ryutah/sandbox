package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
	"github.com/ryutah/sandbox/logging-and-alerting/infra/local"
)

func main() {
	handler := laa.CreateHandler(laa.Injection{
		Logger:         local.NewLogger(),
		Alerter:        local.NewAlerter(),
		ContextFactory: local.NewContextFactory(),
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}
