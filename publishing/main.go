package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

func main() {
	projectId := os.Getenv("GCP_PROJECT_ID")
	auth_json_path := os.Getenv("AUTH_JSON_PATH")
	topic_name := os.Getenv("TOPIC_NAME")
	configOptions := option.WithCredentialsFile(auth_json_path)
	client, err := pubsub.NewClient(context.Background(), projectId, configOptions)
	if err != nil {
		log.Fatal(err)
	}
	topic := client.Topic(topic_name)
	res := topic.Publish(context.Background(), &pubsub.Message{
		Data: []byte("on"),
	})
	msgID, err := res.Get(context.Background())
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print(msgID)
	}
}
