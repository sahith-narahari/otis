package alert

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sahith-narahari/otis/config"
	"log"
	"time"
)

func SendAlert() error {
	bot, err := tgbotapi.NewBotAPI(config.NewApp.Token)
	if err != nil {
		fmt.Println("Error starting bot", err.Error())
		return err
	}
	bot.Debug = true

	update := tgbotapi.NewUpdate(0)
	update.Timeout = 60

	chanUpdates, err := bot.GetUpdatesChan(update)
	if err != nil {
		fmt.Printf("unable to get updates")
	}

	for updates := range chanUpdates {
		if updates.Message == nil {
			continue
		}

		if updates.Message.IsCommand() {
			msg := tgbotapi.NewMessage(updates.Message.Chat.ID, "hello")
			switch updates.Message.Command() {
			case "start":
				go func() {
					for {
						if err := GetStatus(); err != nil {
							msgDisplay := "Port\t" + config.NewApp.Port + "\tnot working"
							msg.Text = msgDisplay
							bot.Send(msg)
						}
						time.Sleep(time.Second * 2)
					}
				}()
			default:
				msg.Text = "start the bot by using command start"
			}
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
	return nil
}
