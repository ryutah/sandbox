package operations

import (
	"net/http"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

type Config struct {
	ProjectID   string
	TraceConfig trace.Config
	OnError     func(error)
}

func NewHandler(handler http.Handler, conf Config) (http.Handler, error) {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: conf.ProjectID,
		OnError:   conf.OnError,
	})
	if err != nil {
		return nil, err
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(conf.TraceConfig)
	return &ochttp.Handler{
		Handler:     handler,
		Propagation: new(propagation.HTTPFormat),
	}, nil
}
