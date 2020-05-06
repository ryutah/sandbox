package main

import (
	"log"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
	"github.com/ryutah/sandbox/logging-and-alerting/infra/local"
)

func main() {
	log.Fatal(laa.Start(laa.Injection{
		Logger:         local.NewLogger(),
		Alerter:        local.NewAlerter(),
		ContextFactory: local.NewContextFactory(),
	}))
}
