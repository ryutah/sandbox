module github.com/ryutah/sandbox/cloud-run-helloworld

go 1.13

require (
	cloud.google.com/go v0.45.1
	contrib.go.opencensus.io/exporter/stackdriver v0.13.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/caarlos0/env/v6 v6.1.0
	github.com/google/uuid v1.1.1
	github.com/openzipkin/zipkin-go v0.2.2
	go.opencensus.io v0.22.1
)
