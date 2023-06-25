package bot

import (
	"fmt"
	"log"

	"github.com/po1nt-1/fantastic-happiness/internal/config"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	bot, err := tgBotApi.NewBotAPI(config.Config.Tg.Token)
	if err != nil {
		log.Fatalf("tgBotApi.NewBotAPI: %v", err)
	}
	bot.Debug = config.Config.Tg.Debug
	log.Printf("Authorized on account: %v", bot.Self.UserName)

	u := tgBotApi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		msg := update.Message
		if msg == nil {
			continue
		}
		if msg.Photo == nil {
			continue
		}

		fileID := msg.Photo[len(msg.Photo)-1].FileID
		fileURL, err := bot.GetFileDirectURL(fileID)
		if err != nil {
			log.Printf("bot.GetFileDirectURL: %v", err)
			continue
		}

		responseText, err := processImage(fileURL)
		if err != nil {
			log.Printf("processImage: %v: %v", err, responseText)
			responseText = fmt.Sprintf("server error: %v", responseText)
		}

		responseText = tgBotApi.EscapeText(tgBotApi.ModeMarkdownV2, responseText)
		response := tgBotApi.NewMessage(msg.Chat.ID, responseText)
		response.ReplyToMessageID = msg.MessageID
		response.ParseMode = tgBotApi.ModeMarkdownV2
		if _, err := bot.Send(response); err != nil {
			log.Printf("bot.Send: %v", err)			
		}
	}
}
