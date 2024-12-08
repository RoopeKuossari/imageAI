package main

import (
	"log"
	"net/http"
	"os"

	"omat/imageAI/env"
	"omat/imageAI/telegram"
)

func main() {
	// Load environment variables
	env.LoadEnv()

	// Fetch bot token from environment
	botToken := os.Getenv("TELEGRAM_API_KEY")
	if botToken == "" {
		log.Fatal("TELEGRAM_API_KEY environment variable is not set")
	}

	// Initialize Telegram bot
	bot, err := telegram.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Start Telegram updates channel
	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	// Start HTTP server
	http.HandleFunc("/hello", telegram.HandleHelloCommand(bot))
	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("HTTP server failed:", err)
		}
	}()

	// Process updates from Telegram
	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if update.Message.Text == "/hello" {
				msg := telegram.NewMessage(update.Message.Chat.ID, "Hello!")
				if _, err := bot.Send(msg); err != nil {
					log.Println("Error sending message:", err)
				}
			}
		}
	}
}
