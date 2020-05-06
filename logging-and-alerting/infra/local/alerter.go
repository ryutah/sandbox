package local

import (
	"context"
	"log"

	laa "github.com/ryutah/sandbox/logging-and-alerting"
)

type Alerter struct {
}

var _ laa.Alerter = new(Alerter)

func NewAlerter() *Alerter {
	return &Alerter{}
}

func (a *Alerter) Alert(ctx context.Context, err error) {
	log.Printf("alert error: %v", err)
}
