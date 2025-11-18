package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ZaguanLabs/perplexity-go/perplexity"
	"github.com/ZaguanLabs/perplexity-go/perplexity/chat"
	"github.com/ZaguanLabs/perplexity-go/perplexity/types"
)

func main() {
	// Create client (reads API key from PERPLEXITY_API_KEY environment variable)
	client, err := perplexity.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	// Create a simple chat completion
	params := &chat.CompletionParams{
		Model: "sonar",
		Messages: []types.ChatMessage{
			types.SystemMessage("You are a helpful assistant."),
			types.UserMessage("What is the capital of France?"),
		},
		MaxTokens:   types.Int(100),
		Temperature: types.Float64(0.7),
	}

	fmt.Println("Sending chat completion request...")

	result, err := client.Chat.Create(context.Background(), params)
	if err != nil {
		log.Fatalf("Chat completion failed: %v", err)
	}

	// Print the response
	fmt.Printf("\nResponse ID: %s\n", result.ID)
	fmt.Printf("Model: %s\n", result.Model)

	if len(result.Choices) > 0 {
		choice := result.Choices[0]
		fmt.Printf("\nAssistant: ")

		// Handle both text and structured content
		switch content := choice.Message.Content.(type) {
		case types.TextContent:
			fmt.Println(string(content))
		case types.StructuredContent:
			for _, chunk := range content {
				if textChunk, ok := chunk.(types.TextChunk); ok {
					fmt.Println(textChunk.Text)
				}
			}
		}

		if choice.FinishReason != nil {
			fmt.Printf("Finish reason: %s\n", *choice.FinishReason)
		}
	}

	// Print usage information
	if result.Usage != nil {
		fmt.Printf("\nUsage:\n")
		fmt.Printf("  Prompt tokens: %d\n", result.Usage.PromptTokens)
		fmt.Printf("  Completion tokens: %d\n", result.Usage.CompletionTokens)
		fmt.Printf("  Total tokens: %d\n", result.Usage.TotalTokens)
		fmt.Printf("  Total cost: $%.6f\n", result.Usage.Cost.TotalCost)
	}

	// Print citations if available
	if len(result.Citations) > 0 {
		fmt.Printf("\nCitations:\n")
		for i, citation := range result.Citations {
			fmt.Printf("  [%d] %s\n", i+1, citation)
		}
	}

	// Print search results if available
	if len(result.SearchResults) > 0 {
		fmt.Printf("\nSearch Results:\n")
		for i, sr := range result.SearchResults {
			fmt.Printf("  [%d] %s\n", i+1, sr.Title)
			fmt.Printf("      %s\n", sr.URL)
		}
	}

	// Example with web search options
	fmt.Println("\n" + string(make([]byte, 50)) + "\n")
	fmt.Println("Example with web search:")

	searchParams := &chat.CompletionParams{
		Model: "sonar",
		Messages: []types.ChatMessage{
			types.UserMessage("What are the latest developments in AI?"),
		},
		MaxTokens:              types.Int(200),
		ReturnRelatedQuestions: types.Bool(true),
		SearchRecencyFilter:    (*chat.SearchRecencyFilter)(types.String(string(chat.SearchRecencyWeek))),
	}

	fmt.Println("Note: Uncomment the code below to run with a real API key")
	_ = searchParams // Prevent unused variable error

	/*
		searchResult, err := client.Chat.Create(context.Background(), searchParams)
		if err != nil {
			log.Fatalf("Search-enabled chat failed: %v", err)
		}

		fmt.Printf("Response: %v\n", searchResult)
	*/

	// Check if API key is set
	if os.Getenv("PERPLEXITY_API_KEY") == "" {
		fmt.Println("\n⚠️  Set PERPLEXITY_API_KEY environment variable to run this example")
	}
}
