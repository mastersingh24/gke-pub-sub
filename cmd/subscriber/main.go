/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"context"
	"flag"
	"log"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/pubsub"
)

func main() {
	// commandline flags
	subID := flag.String("subscription_id", "", "Name of Cloud Pub/Sub subscription")

	// parse flags
	flag.Parse()

	hostname, err := metadata.Hostname()
	if err != nil {
		log.Fatalf("Failed to get hostname: %v\n", err)
	}
	log.Printf("hostname: %s\n", hostname)

	// get the GCP ProjectID
	projectID, err := metadata.ProjectID()

	if err != nil {
		log.Fatalf("Failed to get GCP ProjectID: %v\n", err)
	}

	log.Printf("ProjectID: %s\n", projectID)
	log.Printf("Creating pubsub client for subscription: %s\n", *subID)

	// create pubsub client
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create pubsub client: %v\n", err)
	}

	// subscribe to subID
	sub := client.Subscription(*subID)
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.Printf("Received message with ID [%s]: %q\n", msg.ID, string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error receiving messages: %v\n", err)
	}
}
