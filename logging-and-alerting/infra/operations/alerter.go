package operations

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"cloud.google.com/go/logging"
	laa "github.com/ryutah/sandbox/logging-and-alerting"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
	logpb "google.golang.org/genproto/googleapis/logging/v2"
)

type Alerter struct {
	projectID string
	service   string
	version   string
	logger    *logging.Logger
}

var _ laa.Alerter = new(Alerter)

func NewAlerter(projectID, service, version string, r ResourceConfig) *Alerter {
	ctx := context.Background()
	client, err := logging.NewClient(ctx, fmt.Sprintf("projects/%s", projectID))
	if err != nil {
		panic(err)
	}
	return &Alerter{
		projectID: projectID,
		service:   service,
		version:   version,
		logger: client.Logger("alert_logs", logging.CommonResource(&mrpb.MonitoredResource{
			Type:   r.Type,
			Labels: r.Labels,
		})),
	}
}

func (a *Alerter) Alert(ctx context.Context, err error) {
	dump := fmt.Sprintf("%+v", err)
	dumps := strings.Split(dump, "\n")
	causeFunc, causeFileLine := dumps[1], strings.Split(dumps[2], ":")
	causeLine, _ := strconv.Atoi(strings.TrimSpace(causeFileLine[1]))
	pc, file, line, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	req := extractRequest(ctx)

	a.logger.Log(logging.Entry{
		Severity: logging.Alert,
		Payload: map[string]interface{}{
			"@type":   "type.googleapis.com/google.devtools.clouderrorreporting.v1beta1.ReportedErrorEvent",
			"message": err.Error(),
			"context": map[string]interface{}{
				"httpRequest": map[string]interface{}{
					"method":    req.request.Method,
					"url":       req.request.URL.String(),
					"referrer":  req.request.Referer(),
					"userAgent": req.request.UserAgent(),
				},
				"reportLocation": map[string]interface{}{
					"filePath":     strings.TrimSpace(causeFileLine[0]),
					"lineNumber":   causeLine,
					"functionName": strings.TrimSpace(causeFunc),
				},
			},
			"serviceContext": map[string]string{
				"service": a.service,
				"version": a.version,
			},
		},
		Trace: fmt.Sprintf("projects/%s/traces/%s", a.projectID, extractTraceID(ctx)),
		SourceLocation: &logpb.LogEntrySourceLocation{
			File:     file,
			Line:     int64(line),
			Function: f.Name(),
		},
	})
}
