package operations

import (
	"context"
	"fmt"
	"runtime"

	"cloud.google.com/go/logging"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

type ResourceConfig struct {
	Labels map[string]string
	Type   string
}

type Logger struct {
	projectID string
	logger    *logging.Logger
}

var _ laa.Logger = new(Logger)

func NewLogger(projectID string, r ResourceConfig) *Logger {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, fmt.Sprintf("projects/%s", projectID))
	if err != nil {
		panic(err)
	}
	return &Logger{
		projectID: projectID,
		logger: client.Logger("app_logs", logging.CommonResource(&mrpb.MonitoredResource{
			Type:   r.Type,
			Labels: r.Labels,
		})),
	}
}

func (l *Logger) Infof(ctx context.Context, format string, v ...interface{}) {
	l.printf(ctx, logging.Info, format, v...)
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...interface{}) {
	l.printf(ctx, logging.Error, format, v...)
}

func (l *Logger) printf(ctx context.Context, s logging.Severity, format string, v ...interface{}) {
	pc, file, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	l.logger.Log(logging.Entry{
		Severity: s,
		Payload:  fmt.Sprintf(format, v...),
		Trace:    fmt.Sprintf("projects/%s/traces/%s", l.projectID, extractTraceID(ctx)),
		SourceLocation: &logpb.LogEntrySourceLocation{
			File:     file,
			Line:     int64(line),
			Function: f.Name(),
		},
	})
}
