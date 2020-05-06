package main

import (
	"log"
	"os"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
	"github.com/ryutah/sandbox/logging-and-alerting/infra/operations"
)

func main() {
	var (
		projectID      = os.Getenv("GOOGLE_CLOUD_PROJECT")
		service        = os.Getenv("GAE_SERVICE")
		version        = os.Getenv("GAE_VERSION")
		resourceConfig = operations.ResourceConfig{
			Labels: map[string]string{
				"module_id":  service,
				"version_id": version,
				"project_id": projectID,
			},
			Type: "gae_app",
		}
	)
	log.Fatal(laa.Start(laa.Injection{
		Logger:         operations.NewLogger(projectID, resourceConfig),
		Alerter:        operations.NewAlerter(projectID, service, version, resourceConfig),
		ContextFactory: operations.NewContextFactory(),
	}))
}
