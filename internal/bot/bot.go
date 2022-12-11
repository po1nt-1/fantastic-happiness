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
		if update.Message == nil {
			continue
		}

		msg := tgBotApi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		if len(update.Message.Text) == 0 {
			continue
		}
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Fatalln(err)
		}
	}
}
