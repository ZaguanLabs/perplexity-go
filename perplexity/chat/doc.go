// Package chat provides access to Perplexity's chat completion APIs.
//
// This package implements both standard and streaming chat completions,
// supporting all features of the Perplexity API including tool calling,
// reasoning traces, and web search integration.
//
// # Basic Usage
//
// Create a simple chat completion:
//
//	result, err := client.Chat.Create(ctx, &chat.CompletionParams{
//		Model: "sonar",
//		Messages: []types.ChatMessage{
//			types.UserMessage("What is the weather like today?"),
//		},
//		MaxTokens: types.Int(100),
//	})
//
// # Streaming
//
// For real-time responses, use streaming:
//
//	stream, err := client.Chat.CreateStream(ctx, &chat.CompletionParams{
//		Model: "sonar",
//		Messages: []types.ChatMessage{
//			types.UserMessage("Tell me a story"),
//		},
//	})
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer stream.Close()
//
//	for stream.Next() {
//		chunk := stream.Current()
//		if len(chunk.Choices) > 0 {
//			fmt.Print(chunk.Choices[0].Delta.Content)
//		}
//	}
//
// # Web Search
//
// Enable web search for up-to-date information:
//
//	result, err := client.Chat.Create(ctx, &chat.CompletionParams{
//		Model: "sonar",
//		Messages: []types.ChatMessage{
//			types.UserMessage("What are the latest AI developments?"),
//		},
//		SearchMode:             types.SearchModePtr(types.SearchModeWeb),
//		SearchRecencyFilter:    types.SearchRecencyFilterPtr(types.SearchRecencyDay),
//		ReturnRelatedQuestions: types.Bool(true),
//		ReturnImages:           types.Bool(true),
//	})
//
// # Tool Calling
//
// Define and use tools (functions) in your chat:
//
//	tools := []types.Tool{
//		{
//			Type: types.ToolTypeFunction,
//			Function: types.ToolFunction{
//				Name:        "get_weather",
//				Description: types.String("Get the current weather"),
//				Parameters: map[string]interface{}{
//					"type": "object",
//					"properties": map[string]interface{}{
//						"location": map[string]interface{}{
//							"type": "string",
//							"description": "City name",
//						},
//					},
//				},
//			},
//		},
//	}
//
//	result, err := client.Chat.Create(ctx, &chat.CompletionParams{
//		Model:    "sonar",
//		Messages: messages,
//		Tools:    tools,
//	})
package chat
