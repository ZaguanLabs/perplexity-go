package main

import (
	"fmt"
	"log"

	"github.com/perplexityai/perplexity-go/perplexity"
	"github.com/perplexityai/perplexity-go/perplexity/types"
)

func main() {
	// Create a new client
	// API key can be provided directly or via PERPLEXITY_API_KEY environment variable
	client, err := perplexity.NewClient(
		"",
		perplexity.WithTimeout(30),
		perplexity.WithMaxRetries(3),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Perplexity Go SDK v%s\n", client.Version())
	fmt.Printf("Client initialized with base URL: %s\n", client.BaseURL())

	// Example: Create messages using helper functions
	messages := []types.ChatMessage{
		types.SystemMessage("You are a helpful assistant."),
		types.UserMessage("Hello, world!"),
	}

	fmt.Printf("Created %d messages\n", len(messages))

	// Example: Use pointer helpers for optional fields
	maxTokens := types.Int(100)
	temperature := types.Float64(0.7)

	fmt.Printf("Max tokens: %d, Temperature: %.1f\n", *maxTokens, *temperature)

	// API methods will be available in upcoming releases
	fmt.Println("\nNote: Chat and Search APIs are coming in Phase 3 & 5")
}
