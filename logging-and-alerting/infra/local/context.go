package local

import (
	"context"
	"net/http"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
)

type ContextFactory struct{}

var _ laa.ContextFactory = new(ContextFactory)

func NewContextFactory() *ContextFactory {
	return &ContextFactory{}
}

func (c *ContextFactory) New(r *http.Request) context.Context {
	return r.Context()
}
