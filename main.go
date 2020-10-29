package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os/exec"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TELEGRAM KEY")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "front" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Start front update")
			bot.Send(msg)
			if err = exec.Command("/usr/bin/git", "-C", "/var/www/papka", "pull").Run(); err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("err with front update, err: %s", err))
				bot.Send(msg)
				continue
			}
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Finish front update")
			bot.Send(msg)
		}
	}
}
