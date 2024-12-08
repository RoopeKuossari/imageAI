package telegram

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Create a new Telegram bot
func NewBotAPI(botToken string) (*tgbotapi.BotAPI, error) {
	return tgbotapi.NewBotAPI(botToken)
}

// Create a new update configuration
func NewUpdate(offset int) tgbotapi.UpdateConfig {
	u := tgbotapi.NewUpdate(offset)
	u.Timeout = 60
	return u
}

// HTTP handler for the /hello endpoint
func HandleHelloCommand(bot *tgbotapi.BotAPI) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chatID := os.Getenv("TELEGRAM_CHAT_ID")
		if chatID == "" {
			log.Println("TELEGRAM_CHAT_ID environment variable is not set")
			http.Error(w, "TELEGRAM_CHAT_ID not set", http.StatusInternalServerError)
			return
		}

		msg := tgbotapi.NewMessageToChannel(chatID, "Hello from HTTP!")
		_, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			http.Error(w, "Failed to send message", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("OK"))
	}
}

// Create a new message configuration
func NewMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}
