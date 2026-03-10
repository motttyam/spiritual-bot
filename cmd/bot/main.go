package main

import (
	"context"
	"log"
	"os"

	"github.com/motttyam/spiritual-bot/internal/fortune"
	"github.com/motttyam/spiritual-bot/internal/mixi2"
)

func main() {
	cfg := mixi2.Config{
		ClientID:     mustGetEnv("MIXI2_CLIENT_ID"),
		ClientSecret: mustGetEnv("MIXI2_CLIENT_SECRET"),
		TokenURL:     mustGetEnv("MIXI2_TOKEN_URL"),
		APIAddress:   mustGetEnv("MIXI2_API_ADDRESS"),
	}

	generator, err := fortune.NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}

	text, err := generator.Generate()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("posting: %s", text)

	client, err := mixi2.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Post(context.Background(), text); err != nil {
		log.Fatal(err)
	}
}

func mustGetEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("required environment variable is not set: %s", key)
	}
	return v
}
