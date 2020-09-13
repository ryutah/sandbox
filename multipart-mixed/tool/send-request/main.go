package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"strings"
)

func main() {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	for i := 0; i < 10; i++ {
		part, err := w.CreatePart(textproto.MIMEHeader{
			"content-type": []string{"application/json"},
			"content-id":   []string{fmt.Sprintf("%d", i+1)},
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
	fmt.Println(string(respDump))

	val := resp.Header.Get("content-type")
	if !strings.HasPrefix(strings.ToLower(val), "multipart/mixed") {
		panic(fmt.Sprintf("content type is %v", val))
	}
	val = strings.TrimSpace(strings.TrimPrefix(val, "multipart/mixed"))
	var boundary string
	for _, v := range strings.Split(val, ";") {
		v = strings.TrimSpace(v)
		fmt.Println(v)
		if strings.HasPrefix(strings.ToLower(v), "boundary") {
			boundary = strings.ReplaceAll(strings.TrimSpace(strings.Split(v, "=")[1]), `"`, "")
		}
	}
	if boundary == "" {
		panic("boundary is blank")
	}

	mr := multipart.NewReader(resp.Body, boundary)
	fmt.Println("================== response ==================")
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		for k, v := range part.Header {
			fmt.Printf("%s: %s\n", k, strings.Join(v, ","))
		}
		body, err := ioutil.ReadAll(part)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(body))
	}
	fmt.Println("=============================================")
}
