package translate

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

func Translate(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		panic(err)
	}

	bodyStrings := strings.Split(string(body), " ")

	sourceLanguage, err := language.Parse(bodyStrings[0])
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read source language: %s", err), http.StatusInternalServerError)
		panic(err)
	}

	// Set target language
	targetLanguage := bodyStrings[1]

	// Sets the text to translate.
	text := strings.Join(bodyStrings[2:], " ")

	ctx := context.Background()

	// Creates a client.
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the target language.
	target, err := language.Parse(targetLanguage)
	if err != nil {
		log.Fatalf("Failed to parse target language: %v", err)
	}

	// Translates the text into Russian.
	translations, err := client.Translate(ctx, []string{text}, target, &translate.Options{Source: sourceLanguage, Format: "text"})
	if err != nil {
		log.Fatalf("Failed to translate text: %v", err)
	}

	//fmt.Printf("Text: %v\n", text)
	//fmt.Printf("Translation: %v\n", translations[0].Text)

	fmt.Fprintf(w, "%v\n", translations[0].Text)

}

func Languages(w http.ResponseWriter, r *http.Request) {

	for _, tag := range display.Supported.Tags() {

		fmt.Fprintf(w, "%v ", tag)
		//fmt.Printf("%v ", tag)
	}
}

// DO NOT FORGET:
// $ export GCP_PROJECT="cloud-functions-talk-22365"
// $ export GOOGLE_APPLICATION_CREDENTIALS="/Users/stefan/.secret/cloud-functions-talk-22365-6ba5c5af57f4.json"

// cd cd ~/go/src/hello-world/services
// gcloud alpha functions deploy translate --region europe-west1 --entry-point Translate --runtime go111 --trigger-http

// curl -d "en ru Hello, world"  https://europe-west1-cloud-functions-talk-22365.cloudfunctions.net/translate
// gcloud alpha functions deploy languages  --region europe-west1 --entry-point Languages  --runtime go111 --trigger-http
//
