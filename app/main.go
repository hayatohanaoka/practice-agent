package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"google.golang.org/genai"

	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/model/gemini"
	"google.golang.org/adk/v2/server/adkrest"
	"google.golang.org/adk/v2/session"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/geminitool"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	timeAgent, err := newTimeAgent(model, []tool.Tool{geminitool.GoogleSearch{}})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	restServer, err := adkrest.NewServer(adkrest.ServerConfig{
		AgentLoader:     agent.NewSingleLoader(timeAgent),
		SessionService:  session.InMemoryService(),
		SSEWriteTimeout: 120 * time.Second,
	})
	if err != nil {
		log.Fatalf("Failed to create REST API server: %v", err)
	}

	mux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", restServer))

	mux.HandleFunc("/v1/systems/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("pong")); err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	})

	log.Println("Starting server on :8080")
	log.Println("API available at http://localhost:8080/api/")
	log.Println("Health check at http://localhost:8080/health")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
