package main

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	log.Println("🔁 Loading .env.test")

	err := godotenv.Load(".env.test")
	if err != nil {
		log.Fatalf("❌ Failed to load .env.test: %v", err)
	}

	log.Println("✅ .env.test loaded, DB_STRING:", os.Getenv("DB_STRING"))

	code := m.Run()
	os.Exit(code)
}
