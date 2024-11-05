package t_bot

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
)

type Bot struct {
	TgBot *telebot.Bot
	Chat  *telebot.Chat
}
func (b *Bot) Update(message string) {
	_, err := b.TgBot.Send(b.Chat, message)
	if err != nil {
		log.Errorf("Error while sending message in observe: %v", err)
	}
}
