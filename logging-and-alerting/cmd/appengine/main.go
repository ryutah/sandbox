package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/profiler"
	"go.opencensus.io/trace"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
	"github.com/ryutah/sandbox/logging-and-alerting/infra/operations"
	opencensusOperations "github.com/ryutah/sandbox/logging-and-alerting/lib/opencensus/operations"
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

	if err := profiler.Start(profiler.Config{}); err != nil {
		panic(err)
	}
	handler, err := opencensusOperations.NewHandler(
		laa.CreateHandler(laa.Injection{
			Logger:         operations.NewLogger(projectID, resourceConfig),
			Alerter:        operations.NewAlerter(projectID, service, version, resourceConfig),
			ContextFactory: operations.NewContextFactory(),
		}),
		opencensusOperations.Config{
			ProjectID:   projectID,
			OnError:     func(err error) { log.Printf("failed to output trace: %v", err) },
			TraceConfig: trace.Config{DefaultSampler: trace.AlwaysSample()},
		},
	)
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}
