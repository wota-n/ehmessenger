package main

import (
	"context"
	"fmt"
	"log"
	"time"

	eventhub "github.com/Azure/azure-event-hubs-go/v3"
)

func main() {
	// Set up environment variables
	eventHubConnectionString := ""

	// Create a new Event Hub client
	hub, err := eventhub.NewHubFromConnectionString(eventHubConnectionString)

	if err != nil {
		log.Fatal(err)
	}
	defer hub.Close(context.Background())

	// Send some events to the Event Hub
	for i := 0; i < 1000; i++ {
		message := fmt.Sprintf("Event #%d at %s", i+1, time.Now().Format(time.RFC3339))
		if err := hub.Send(context.Background(), eventhub.NewEventFromString(message)); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent event: %s\n", message)
	}
}
