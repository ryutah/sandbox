package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/google/uuid"
	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
	helloworld "github.com/ryutah/sandbox/cloud-run-helloworld"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"contrib.go.opencensus.io/exporter/zipkin"
)

func main() {
	initTracing()
	initProfiler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.Handle("/", helloworld.NewServeMux(traceID))
	mux.HandleFunc("/_healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), &ochttp.Handler{
		Handler: mux,
	}))
}

func initTracing() {
	endpoint, err := openzipkin.NewEndpoint("cloud-run-started", "localhost:8080")
	if err != nil {
		panic(err)
	}
	zipkinHost := os.Getenv("ZIPKIN_HOST")
	reporter := zipkinHTTP.NewReporter(fmt.Sprintf("%s/api/v2/spans", zipkinHost))
	ze := zipkin.NewExporter(reporter, endpoint)
	trace.RegisterExporter(ze)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func initProfiler() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
}

func traceID(r *http.Request) string {
	return uuid.Must(uuid.NewRandom()).String()
}
