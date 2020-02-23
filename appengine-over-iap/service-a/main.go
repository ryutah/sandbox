package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	secretmanager "cloud.google.com/go/secretmanager/apiv1beta1"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1beta1"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"

	"bytes"

	"context"

	"strings"

	"github.com/ryutah/sandbox/appengine-over-iap/lib"
)

var configs = struct {
	projectID   string
	location    string
	queueName   string
	iapClientID string
}{
	projectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	location:  os.Getenv("QUEUE_LOCATION"),
	queueName: "iap-example",
}

var clients = struct {
	tasks *cloudtasks.Client
}{}

func main() {
	var err error
	clients.tasks, err = cloudtasks.NewClient(context.Background())
	if err != nil {
		panic(err)
	}
	defer clients.tasks.Close()

	secretClient, err := secretmanager.NewClient(context.Background())
	if err != nil {
		panic(err)
	}
	defer secretClient.Close()
	secretResp, err := secretClient.AccessSecretVersion(context.Background(), &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/%s/secrets/appeingine-iap-client/versions/latest", configs.projectID),
	})
	if err != nil {
		panic(err)
	}
	configs.iapClientID = strings.TrimSpace(string(secretResp.Payload.Data))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := newClient().Get(fmt.Sprintf("https://service-b-dot-%s.appspot.com/", configs.projectID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Receive: %s", body)
	})

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		_, err := clients.tasks.CreateTask(r.Context(), &tasks.CreateTaskRequest{
			Parent: fmt.Sprintf("projects/%s/locations/%s/queues/%s", configs.projectID, configs.location, configs.queueName),
			Task: &tasks.Task{
				MessageType: &tasks.Task_AppEngineHttpRequest{
					AppEngineHttpRequest: &tasks.AppEngineHttpRequest{
						HttpMethod: tasks.HttpMethod_GET,
						AppEngineRouting: &tasks.AppEngineRouting{
							Service: "service-b",
						},
						RelativeUri: "/",
					},
				},
			},
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("send success"))
	})

	lib.Serve()
}

type roundTripper struct{}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	url := fmt.Sprintf(
		"http://metadata/computeMetadata/v1/instance/service-accounts/default/identity?audience=%s",
		configs.iapClientID,
	)
	idTokenReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	idTokenReq.Header.Set("Metadata-Flavor", "Google")
	idTokenResp, err := http.DefaultClient.Do(idTokenReq)
	if err != nil {
		return nil, err
	}
	defer idTokenResp.Body.Close()
	token, err := ioutil.ReadAll(idTokenResp.Body)
	if err != nil {
		return nil, err
	}

	newReq := *req
	newReq.Body = ioutil.NopCloser(new(bytes.Buffer))

	newReq.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return http.DefaultClient.Do(&newReq)
}

func newClient() *http.Client {
	return &http.Client{
		Transport: new(roundTripper),
	}
}
