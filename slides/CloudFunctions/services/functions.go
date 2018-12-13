package services

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/api/iterator"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
)

// Service is the struct for the collection
type Service struct {
	Name        string `firestore:"name,omitempty"`
	Url         string `firestore:"url,omitempty"`
	Description string `firestore:"description,omitempty"`
}

// Structure of the collection
var collectionName string = "Services"

// Register information of the new service via http and stores it in Firestore
func Register(w http.ResponseWriter, r *http.Request) {

	// Get rid of warnings
	_ = r

	// Sets your Google Cloud Platform project ID.
	projectId := os.Getenv("GCP_PROJECT")
	if projectId == "" {
		http.Error(w, fmt.Sprintf("GCP_PROJECT environment variable unset or missing"), http.StatusInternalServerError)
	}

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %v", err), http.StatusInternalServerError)
	}
	defer client.Close()

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		panic(err)
	}

	// Create a struct out of the body (except first word, i.e. UUID as key/document)
	words := strings.Split(string(body), " ")
	service := &Service{
		Name:        words[0],
		Url:         words[1],
		Description: strings.Join(words[2:], " "),
	}

	// Marshall the struct to JSON
	ipAddressJson, err := json.MarshalIndent(service, "", "  ")
	if err != nil {
		// Return HTTP error code 500 Internal Server Error
		http.Error(w, fmt.Sprintf("failed to marshall %v: %s", service, err), http.StatusInternalServerError)
	}

	// Save the JSON as string in filed named by the UUID and as document named by the same UUID
	_, err = client.Collection(collectionName).Doc(strings.Split(string(body), " ")[0]).Set(ctx, map[string]interface{}{
		words[0]: fmt.Sprintf("%v", string(ipAddressJson)),
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to register %v: %s", service, err), http.StatusInternalServerError)
	}
}

// Unregister deletes the service specified by name in Firestore
func Unregister(w http.ResponseWriter, r *http.Request) {

	// Sets your Google Cloud Platform project ID.
	projectId := os.Getenv("GCP_PROJECT")
	if projectId == "" {
		http.Error(w, fmt.Sprintf("GCP_PROJECT environment variable unset or missing"), http.StatusInternalServerError)
	}

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %v", err), http.StatusInternalServerError)
	}
	defer client.Close()

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		panic(err)
	}

	// Delete document specified by UUID
	_, err = client.Collection(collectionName).Doc(strings.Split(string(body), " ")[0]).Delete(ctx)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to delete %q: %s", strings.Split(string(body), " ")[0], err), http.StatusInternalServerError)
	}
}

// List returns the list of services from Firestore
func List(w http.ResponseWriter, r *http.Request) {

	// Get rid of warnings
	_ = r

	// Sets your Google Cloud Platform project ID.
	projectId := os.Getenv("GCP_PROJECT")
	if projectId == "" {
		http.Error(w, fmt.Sprintf("GCP_PROJECT environment variable unset or missing"), http.StatusInternalServerError)
	}

	// Get a Firestore client.
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create client: %v", err), http.StatusInternalServerError)
	}
	defer client.Close()

	iter := client.Collection(collectionName).Documents(ctx)

	// Map of Service
	var services map[string]Service = make(map[string]Service)

	// Iterate over the documents
	var service Service
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to iterate over collection %q: %s", collectionName, err), http.StatusInternalServerError)
		}

		// Get the JSON string, unmarshall it, and insert it into the map
		if v, ok := doc.Data()[doc.Ref.ID].(string); ok {
			json.Unmarshal([]byte(v), &service)
		} else {
			http.Error(w, fmt.Sprintf("failed to convert %q: %v", collectionName, err), http.StatusInternalServerError)
		}
		services[doc.Ref.ID] = service
	}

	// Marshall all and return it as response
	ipAddressesJson, err := json.Marshal(services)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to marshall %q: %v", collectionName, err), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%v", string(ipAddressesJson))
}

// DO NOT FORGET:
// $ export GCP_PROJECT="cloud-functions-talk-22365"
// $ export GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/.secret/cloud-functions-talk-22365-6ba5c5af57f4.json"

// https://console.cloud.google.com/firestore/welcome?project=cloud-functions-talk-22365
// gcloud services enable firestore.googleapis.com

// cd ~/go/src/hello-world/services
// gcloud alpha functions deploy register --region europe-west1 --entry-point Register --runtime go111 --trigger-http

// curl -d "languages https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/languages languages shows all known language tags, e.g. 'service languages'"  https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/register

// gcloud alpha functions deploy unregister  --region europe-west1 --entry-point Unregister  --runtime go111 --trigger-http

// curl -d "languages"  https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/unregister

// gcloud alpha functions deploy services  --region europe-west1 --entry-point List  --runtime go111 --trigger-http

// curl -d "translate https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/translate translate from one language to another, e.g. 'service translate en ru hello world'"  https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/register

// curl -d "en ru Hello, world"  https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/translate

// curl -d "myip https://europe-west1-gke-serverless-211907.cloudfunctions.net/myip myip shows my IP address as seen from the server, e.g. 'service myip'" https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/register
