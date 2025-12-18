// Package perplexity provides a Go client library for the Perplexity API.
//
// This package offers idiomatic Go interfaces for interacting with Perplexity's
// chat completions, streaming responses, and web search capabilities.
//
// # Quick Start
//
// Create a new client and make a chat completion request:
//
//	client, err := perplexity.NewClient("your-api-key")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	result, err := client.Chat.Create(context.Background(), &chat.CompletionParams{
//		Model: "sonar",
//		Messages: []types.ChatMessage{
//			types.UserMessage("What is the capital of France?"),
//		},
//	})
//
// # Configuration
//
// The client can be configured with functional options:
//
//	client, err := perplexity.NewClient(
//		"your-api-key",
//		perplexity.WithTimeout(30*time.Second),
//		perplexity.WithMaxRetries(5),
//		perplexity.WithBaseURL("https://custom.api.com"),
//	)
//
// # Error Handling
//
// The SDK provides typed errors for different HTTP status codes:
//
//	resp, err := client.Chat.Create(ctx, params)
//	if err != nil {
//		switch e := err.(type) {
//		case *perplexity.AuthenticationError:
//			// Handle authentication error
//		case *perplexity.RateLimitError:
//			// Handle rate limit error
//		default:
//			// Handle other errors
//		}
//	}
//
// # Streaming
//
// For streaming responses, use CreateStream:
//
//	stream, err := client.Chat.CreateStream(ctx, params)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stream.Close()
//
//	for stream.Next() {
//		chunk := stream.Current()
//		fmt.Print(chunk.Choices[0].Delta.Content)
//	}
//
// # API Key
//
// The API key can be provided in two ways:
//  1. Directly as a parameter to NewClient
//  2. Via the PERPLEXITY_API_KEY environment variable
//
// For more information, see https://docs.perplexity.ai/
package perplexity
