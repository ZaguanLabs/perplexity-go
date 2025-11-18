package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/perplexityai/perplexity-go/perplexity"
	"github.com/perplexityai/perplexity-go/perplexity/chat"
	"github.com/perplexityai/perplexity-go/perplexity/types"
)

func main() {
	// Create client (reads API key from PERPLEXITY_API_KEY environment variable)
	client, err := perplexity.NewClient("")
	if err != nil {
		log.Fatal(err)
	}

	// Example 1: Streaming with Next()
	fmt.Println("=== Example 1: Streaming with Next() ===")
	fmt.Println()
	streamWithNext(client)

	fmt.Println()
	fmt.Println(string(make([]byte, 50)))
	fmt.Println()

	// Example 2: Streaming with Iter() channel
	fmt.Println("=== Example 2: Streaming with Iter() ===")
	fmt.Println()
	streamWithIter(client)

	// Check if API key is set
	if os.Getenv("PERPLEXITY_API_KEY") == "" {
		fmt.Println("\n⚠️  Set PERPLEXITY_API_KEY environment variable to run this example")
	}
}

func streamWithNext(client *perplexity.Client) {
	params := &chat.CompletionParams{
		Model: "sonar",
		Messages: []types.ChatMessage{
			types.UserMessage("Write a haiku about programming"),
		},
		MaxTokens:   types.Int(100),
		Temperature: types.Float64(0.7),
	}

	fmt.Println("Creating streaming completion...")
	stream, err := client.Chat.CreateStream(context.Background(), params)
	if err != nil {
		log.Printf("CreateStream failed: %v", err)
		return
	}
	defer stream.Close()

	fmt.Print("Response: ")

	var fullText string
	for {
		chunk, err := stream.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Stream error: %v", err)
			return
		}

		// Print each chunk as it arrives
		if len(chunk.Choices) > 0 {
			delta := chunk.Choices[0].Delta
			switch content := delta.Content.(type) {
			case types.TextContent:
				text := string(content)
				fmt.Print(text)
				fullText += text
			case types.StructuredContent:
				for _, c := range content {
					if textChunk, ok := c.(types.TextChunk); ok {
						fmt.Print(textChunk.Text)
						fullText += textChunk.Text
					}
				}
			}
		}

		// Print usage info on final chunk
		if chunk.Usage != nil {
			fmt.Printf("\n\nUsage:\n")
			fmt.Printf("  Prompt tokens: %d\n", chunk.Usage.PromptTokens)
			fmt.Printf("  Completion tokens: %d\n", chunk.Usage.CompletionTokens)
			fmt.Printf("  Total tokens: %d\n", chunk.Usage.TotalTokens)
			fmt.Printf("  Total cost: $%.6f\n", chunk.Usage.Cost.TotalCost)
		}
	}

	fmt.Println()
}

func streamWithIter(client *perplexity.Client) {
	params := &chat.CompletionParams{
		Model: "sonar",
		Messages: []types.ChatMessage{
			types.UserMessage("Count from 1 to 5"),
		},
		MaxTokens: types.Int(50),
	}

	fmt.Println("Creating streaming completion...")
	stream, err := client.Chat.CreateStream(context.Background(), params)
	if err != nil {
		log.Printf("CreateStream failed: %v", err)
		return
	}
	defer stream.Close()

	fmt.Print("Response: ")

	// Use channel-based iteration
	for chunk := range stream.Iter() {
		if len(chunk.Choices) > 0 {
			delta := chunk.Choices[0].Delta
			switch content := delta.Content.(type) {
			case types.TextContent:
				fmt.Print(string(content))
			case types.StructuredContent:
				for _, c := range content {
					if textChunk, ok := c.(types.TextChunk); ok {
						fmt.Print(textChunk.Text)
					}
				}
			}
		}
	}

	fmt.Println()
}

// Example with context cancellation
func streamWithCancellation() {
	client, _ := perplexity.NewClient("")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*1000) // 5 seconds
	defer cancel()

	params := &chat.CompletionParams{
		Model: "sonar",
		Messages: []types.ChatMessage{
			types.UserMessage("Write a long story"),
		},
		MaxTokens: types.Int(1000),
	}

	stream, err := client.Chat.CreateStream(ctx, params)
	if err != nil {
		log.Printf("CreateStream failed: %v", err)
		return
	}
	defer stream.Close()

	for {
		chunk, err := stream.Next()
		if err == io.EOF {
			fmt.Println("\nStream completed")
			break
		}
		if err == context.DeadlineExceeded {
			fmt.Println("\nStream cancelled due to timeout")
			break
		}
		if err != nil {
			log.Printf("Stream error: %v", err)
			break
		}

		// Process chunk...
		_ = chunk
	}
}
