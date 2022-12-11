package bot

import (
	"log"

	"fantastic-happiness/internal/config"

	tgBotApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run() {
	bot, err := tgBotApi.NewBotAPI(config.Config.Tg.Token)
	if err != nil {
		log.Fatalln(err)
		return
	}
	bot.Debug = true

	log.Println("Authorized on account", bot.Self.UserName)

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
			log.Fatalln(err)
		}
		log.Println(fileURL)

		response := tgBotApi.NewMessage(msg.Chat.ID, "picture is received")
		response.ReplyToMessageID = msg.MessageID
		response.ParseMode = tgBotApi.ModeMarkdownV2

		if _, err := bot.Send(response); err != nil {
			log.Fatalln(err)
		}
	}
}
