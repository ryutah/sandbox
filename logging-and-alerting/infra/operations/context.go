package operations

import (
	"context"
	"net/http"
	"strings"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
)

type ContextFactory struct{}

var _ laa.ContextFactory = new(ContextFactory)

func NewContextFactory() *ContextFactory {
	return &ContextFactory{}
}

func (c *ContextFactory) New(r *http.Request) context.Context {
	trace := r.Header.Get("X-Cloud-Trace-Context")
	traceID := strings.SplitN(trace, "/", 2)[0]
	ctx := context.WithValue(r.Context(), contextKeyTraceID{}, traceID)
	ctx = context.WithValue(ctx, contextKeyRequest{}, &requestInformation{
		request: r,
	})
	return ctx
}

type contextKeyTraceID struct{}

func extractTraceID(ctx context.Context) string {
	val, ok := ctx.Value(contextKeyTraceID{}).(string)
	if !ok {
		return ""
	}
	return val
}

type contextKeyRequest struct{}

type requestInformation struct {
	request *http.Request
}

func extractRequest(ctx context.Context) *requestInformation {
	val, ok := ctx.Value(contextKeyRequest{}).(*requestInformation)
	if !ok {
		return nil
	}
	return val
}
