package helloworld

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type traceIDGen func(*http.Request) string

func NewServeMux(gen traceIDGen) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", newHelloWorld(gen).handle)
	return mux
}

type helloWorld struct {
	genTraceID traceIDGen
}

func newHelloWorld(gen traceIDGen) *helloWorld {
	return &helloWorld{
		genTraceID: gen,
	}
}

func (h *helloWorld) handle(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, "", 0)
	entry, _ := json.Marshal(map[string]string{
		"message":                      "called!!",
		"severity":                     "INFO",
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", Configuration().ProjectID, h.genTraceID(r)),
	})
	logger.Println(string(entry))
	w.Write([]byte("ok!"))
}
