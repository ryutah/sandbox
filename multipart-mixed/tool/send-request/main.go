package main

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
)

func main() {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	for i := 0; i < 10; i++ {
		part, err := w.CreatePart(textproto.MIMEHeader{
			"content-type": []string{"application/json"},
		})
		if err != nil {
			panic(err)
		}
		if _, err := fmt.Fprintf(part, `{"message":"part %v"}`, i+1); err != nil {
			panic(err)
		}
	}
	if err := w.Close(); err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080", buf)
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", fmt.Sprintf("multipart/mixed; boundary=%q", w.Boundary()))
	reqDump, _ := httputil.DumpRequestOut(req, true)
	fmt.Println("================== request ==================")
	fmt.Printf("%s\n", reqDump)
	fmt.Println("=============================================")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	respDump, _ := httputil.DumpResponse(resp, true)
	fmt.Println("================== response ==================")
	fmt.Printf("%s\n", respDump)
	fmt.Println("=============================================")
}
