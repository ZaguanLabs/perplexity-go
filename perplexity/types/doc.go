// Package types provides type definitions for the Perplexity API.
//
// This package contains all the request and response types used by the Perplexity API,
// including chat messages, completions, search results, and usage information.
//
// # Message Types
//
// Chat messages can be created using helper functions:
//
//	messages := []types.ChatMessage{
//		types.SystemMessage("You are a helpful assistant."),
//		types.UserMessage("Hello!"),
//		types.AssistantMessage("Hi! How can I help you?"),
//	}
//
// # Structured Content
//
// Messages can contain structured content with multiple chunks:
//
//	msg := types.ChatMessage{
//		Role: types.RoleUser,
//		Content: types.StructuredContent{
//			types.TextChunk{Type: "text", Text: "What's in this image?"},
//			types.ImageChunk{Type: "image_url", ImageURL: types.ImageURLString("https://example.com/image.jpg")},
//		},
//	}
//
// # Helper Functions
//
// The package provides helper functions for creating pointers to primitive types:
//
//	params := &chat.CompletionParams{
//		Model:       "sonar",
//		Messages:    messages,
//		MaxTokens:   types.Int(100),
//		Temperature: types.Float64(0.7),
//		TopP:        types.Float64(0.9),
//	}
//
// # Response Types
//
// Responses include detailed usage and cost information:
//
//	result, err := client.Chat.Create(ctx, params)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Tokens used: %d\n", result.Usage.TotalTokens)
//	fmt.Printf("Cost: $%.6f\n", result.Usage.Cost.TotalCost)
package types
