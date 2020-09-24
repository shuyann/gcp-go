package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"io"
	"os"
	"sync"
)

func publish(w io.Writer, projectID, topicID, msg string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})
	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("get: %v", err)
	}
	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
	return nil
}

func pullMsgs(w io.Writer, projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}

	var mu sync.Mutex
	received := 0
	sub := client.Subscription(subID)
	cctx, cancel := context.WithCancel(ctx)
	err = sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		mu.Lock()
		defer mu.Unlock()
		fmt.Fprintf(w, "got message: %q\n", string(msg.Data))
		msg.Ack()
		received++
		if received == 10 {
			cancel()
		}
	})
	if err != nil {
		return fmt.Errorf("receive: %v", err)
	}
	return nil
}

func main() {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	topicID := os.Getenv("GOOGLE_CLOUD_PUBSUB_TOPIC")
	subID := os.Getenv("GOOGLE_CLOUD_PUBSUB_SUBSCRIPTION")
	msg := "Hello"
	if err := publish(os.Stdout, projectID, topicID, msg); err != nil {
		panic(err)
	}
	if err := pullMsgs(os.Stdout, projectID, subID); err != nil {
		panic(err)
	}
}
