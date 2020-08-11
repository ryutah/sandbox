package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
)

func main() {
	// Send Mail
	err := smtp.SendMail("localhost:1025", nil, "local-dev@sample.com", []string{"ryutah@sample.com"}, []byte("Hello, World!"))
	if err != nil {
		panic(err)
	}

	resp, err := http.Get("http://localhost:8025/api/v2/messages")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	payload := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		panic(err)
	}
	result, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
}
