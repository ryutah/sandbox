package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
)

func main() {
	funcframework.RegisterHTTPFunction("/", handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}

type computeEngineMetadata map[string]interface{}

type envInformation struct {
	EnvVars               map[string]string   `json:"env_vars"`
	Headers               map[string][]string `json:"headers"`
	ComputeEngineMetadata struct {
		Project  computeEngineMetadata `json:"project"`
		Instance computeEngineMetadata `json:"instance"`
	} `json:"compute_engine_metadata"`
}

func newEnvInformation() *envInformation {
	return &envInformation{
		EnvVars: make(map[string]string),
		Headers: make(map[string][]string),
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	payload := newEnvInformation()
	for _, env := range os.Environ() {
		keyValPair := strings.Split(env, "=")
		payload.EnvVars[keyValPair[0]] = keyValPair[1]
	}
	for k, v := range r.Header {
		payload.Headers[k] = v
	}

	project, err := loadComputeEngineMeatadataRecursive("project")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload.ComputeEngineMetadata.Project = project
	instance, err := loadComputeEngineMeatadataRecursive("instance")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	payload.ComputeEngineMetadata.Instance = instance

	results, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	fmt.Fprintln(w, string(results))
}

func loadComputeEngineMeatadataRecursive(dir string) (computeEngineMetadata, error) {
	uri := fmt.Sprintf("http://metadata.google.internal/computeMetadata/v1/%s/?recursive=true", dir)
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Metadata-Flavor", "Google")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	payload := make(computeEngineMetadata)
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return payload, nil
}
