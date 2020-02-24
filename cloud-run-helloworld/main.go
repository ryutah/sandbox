package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/profiler"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
)

var projectID = os.Getenv("GOOGLE_CLOUD_PROJECT")

func main() {
	initTracing()
	initProfiler()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), &ochttp.Handler{
		Handler:     mux,
		Propagation: new(propagation.HTTPFormat),
	}))
}

func initTracing() {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID:    projectID,
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

func handler(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "", 0)
	entry, _ := json.Marshal(map[string]string{
		"message":                      "called!!",
		"severity":                     "INFO",
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", projectID, traceID(r)),
	})
	logger.Println(string(entry))
	w.Write([]byte("ok!"))
}

func traceID(r *http.Request) string {
	return strings.SplitN(r.Header.Get("X-Cloud-Trace-Context"), "/", 2)[0]
}
