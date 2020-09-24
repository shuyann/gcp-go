package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/errorreporting"
)

var errorClient *errorreporting.Client

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")

	var err error
	errorClient, err = errorreporting.NewClient(ctx, projectID, errorreporting.Config{
		ServiceName: "myservice",
		OnError: func(err error) {
			log.Printf("Could not log error: %v", err)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer errorClient.Close()

	resp, err := http.Get("not-a-valid-url")
	if err != nil {
		logAndPrintError(err)
		return
	}
	log.Print(resp.Status)
}

func logAndPrintError(err error) {
	errorClient.Report(errorreporting.Entry{
		Error: err,
	})
	log.Print(err)
}
