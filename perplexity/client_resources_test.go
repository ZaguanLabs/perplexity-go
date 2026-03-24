package perplexity

import "testing"

func TestNewClient_InitializesNewResources(t *testing.T) {
	client, err := NewClient("test-key")
	if err != nil {
		t.Fatalf("NewClient() error = %v", err)
	}
	if client.Responses == nil {
		t.Fatal("Responses service is nil")
	}
	if client.Embeddings == nil {
		t.Fatal("Embeddings service is nil")
	}
	if client.ContextualizedEmbeddings == nil {
		t.Fatal("ContextualizedEmbeddings service is nil")
	}
	if client.Browser == nil {
		t.Fatal("Browser service is nil")
	}
	if client.Browser.Sessions == nil {
		t.Fatal("Browser.Sessions service is nil")
	}
}
