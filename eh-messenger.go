package main

import (
	"context"
	"fmt"
	"log"
	"time"

	eventhubs "github.com/Azure/azure-event-hubs-go/v3"
)

func main() {
	// Set up environment variables
	eventHubConnectionString := "Endpoint=sb://<your-namespace>.servicebus.windows.net/;SharedAccessKeyName=<your-policy>;SharedAccessKey=<your-policy-key>;EntityPath=<your-event-hub>"
	eventHubName := "<your-event-hub-name>"

	// Create a new Event Hub client
	hub, err := eventhubs.NewHubFromConnectionString(eventHubConnectionString, eventHubName)
	if err != nil {
		log.Fatal(err)
	}
	defer hub.Close(context.Background())

	// Send some events to the Event Hub
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Event #%d at %s", i+1, time.Now().Format(time.RFC3339))
		if err := hub.Send(context.Background(), eventhubs.NewEventFromString(message)); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent event: %s\n", message)
	}
}