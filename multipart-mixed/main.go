package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":8080", nil)
}

type reqPayload struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		return
	}

	respBody := new(bytes.Buffer)
	resp := multipart.NewWriter(respBody)
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "failed to reat part", err.Error())
			return
		}

		var paylod reqPayload
		if err := json.NewDecoder(io.TeeReader(part, os.Stdout)).Decode(&paylod); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "failed to decode json", err.Error())
			return
		}
		respPart, err := resp.CreatePart(textproto.MIMEHeader{
			"content-type": []string{"application/json"},
			"content-id":   []string{part.Header.Get("content-id")},
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "failed create part", err.Error())
			return
		}

		fmt.Fprintf(respPart, `{"message":"hello %s"}`, paylod.Message)
		_ = part.Close()
	}

	if err := resp.Close(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "failed to close multipart writer", err.Error())
		return
	}
	w.Header().Set("content-type", fmt.Sprintf("multipart/mixed; boundary=%q", resp.Boundary()))
	fmt.Fprintln(w, respBody.String())
}
