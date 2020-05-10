package laa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Injection struct {
	Logger
	Alerter
	ContextFactory
}

func CreateHandler(i Injection) http.Handler {
	mux := http.NewServeMux()
	h := &handler{
		logger:         i.Logger,
		alerter:        i.Alerter,
		contextFactory: i.ContextFactory,
	}
	mux.HandleFunc("/", h.Handle)
	return mux
}

type Logger interface {
	Infof(ctx context.Context, format string, v ...interface{})
	Errorf(ctx context.Context, format string, v ...interface{})
}

type Alerter interface {
	Alert(ctx context.Context, err error)
}

type ContextFactory interface {
	New(*http.Request) context.Context
}

type handler struct {
	logger         Logger
	alerter        Alerter
	contextFactory ContextFactory
}

type handleRequestPayload struct {
	Message    string `json:"message"`
	RaiseError bool   `json:"raise_error"`
}

func (h *handler) Handle(w http.ResponseWriter, r *http.Request) {
	var req handleRequestPayload
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx := h.contextFactory.New(r)
	if err := h.execProcess(ctx, req); err != nil {
		h.alerter.Alert(ctx, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "ok")
}

func (h *handler) execProcess(ctx context.Context, p handleRequestPayload) error {
	h.logger.Infof(ctx, "receive message: %v", p.Message)
	if p.RaiseError {
		h.logger.Errorf(ctx, "raise error!")
		return errors.New("sample error")
	}
	return nil
}
