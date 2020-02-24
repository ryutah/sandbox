package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	helloworld "github.com/ryutah/sandbox/cloud-run-helloworld"

	"cloud.google.com/go/profiler"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
)

func main() {
	initTracing()
	initProfiler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), &ochttp.Handler{
		Handler:     helloworld.NewServeMux(traceID),
		Propagation: new(propagation.HTTPFormat),
	}))
}

func initTracing() {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:    helloworld.Configuration().ProjectID,
		MetricPrefix: "cloud-run-started",
		OnError: func(err error) {
			log.Printf("failed to export trace: %v", err)
		},
	})
	if err != nil {
		panic(err)
	}

	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func initProfiler() {
	if err := profiler.Start(profiler.Config{
		Service:        "cloud-run-started",
		ServiceVersion: "1.0.0",
	}); err != nil {
		panic(err)
	}
}

func traceID(r *http.Request) string {
	return strings.SplitN(r.Header.Get("X-Cloud-Trace-Context"), "/", 2)[0]
}
