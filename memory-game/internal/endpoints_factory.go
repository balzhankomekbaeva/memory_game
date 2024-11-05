package t_bot

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	HELP = "⌨️ Home location 🏠 - Permanent Location\n  ➡My location 📍- see your home location\n" +
		"  ➡Add location ➕ - adding home location if there is not exist\n" +
		"  ➡Delete location ✖️ - deleting current home location\n" +
		"  ➡Check location ✔️ - checking your current home location, if are there any crime events\n\n" +
		"⌨ Find crime 🔎 - finding crime events\n" +
		"  ➡My location 🌍 - check your current location\n" +
		"  ➡Send location by map - send location you want\n\n" +
		"⌨ Story 🗃️ - see the story you once looked for\n" +
		"  ➡All stories - see all history\n" +
		"  ➡Clear - clear the history"
)

func NewEndpointsFactory(service UserInfo, ctx context.Context) *endpointsFactory {
	return &endpointsFactory{userRepo: service, ctx: ctx}
}

type endpointsFactory struct {
	userRepo UserInfo
	ctx      context.Context
}

func (ef *endpointsFactory) Start(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromDisk("images/img.png"), Caption: "Hi, " + m.Sender.FirstName + ". Welcome to Memory game bot.\nChoose to continue"}
		b.Send(m.Sender, ">>", &tb.ReplyMarkup{
			InlineKeyboard:      nil,
			ReplyKeyboard:       ReplyKeys,
			ResizeReplyKeyboard: false,
		})
		b.Send(m.Sender, photo)
	}
}

func (ef *endpointsFactory) NewGame(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		b.Send(m.Sender, "Choose the difficulty level", &tb.ReplyMarkup{
			InlineKeyboard:      GameOptionReplyKeys,
			ResizeReplyKeyboard: true,
		})

	}
}

func (ef *endpointsFactory) StartGameWithMedium(b *tb.Bot) func(m *tb.Callback) {
	return func(m *tb.Callback) {
		factory := &DefaultGameFactory{repo: ef.userRepo}
		newGame := factory.NewGame(&MediumDifficulty{})
		chat := m.Message.Chat
		bot := Bot{TgBot: b, Chat: chat}
		b.Send(chat, "Get ready for the memory game!")
		showBoard(bot, chat, *newGame)
		newGame.startGame(&bot, chat)
	}
}

func (ef *endpointsFactory) StartGameWithEasy(b *tb.Bot) func(m *tb.Callback) {
	return func(m *tb.Callback) {
		factory := &DefaultGameFactory{repo: ef.userRepo}
		newGame := factory.NewGame(&EasyDifficulty{})
		chat := m.Message.Chat
		bot := Bot{TgBot: b, Chat: chat}
		b.Send(chat, "Get ready for the memory game!")
		showBoard(bot, chat, *newGame)
		newGame.startGame(&bot, chat)
	}
}

func (ef *endpointsFactory) StartGameWithHard(b *tb.Bot) func(m *tb.Callback) {
	return func(m *tb.Callback) {
		factory := &DefaultGameFactory{repo: ef.userRepo}
		newGame := factory.NewGame(&HardDifficulty{})
		chat := m.Message.Chat
		bot := Bot{TgBot: b, Chat: chat}
		b.Send(chat, "Get ready for the memory game!")
		showBoard(bot, chat, *newGame)
		newGame.startGame(&bot, chat)
	}
}

func (ef *endpointsFactory) StartGameWithNoob(b *tb.Bot) func(m *tb.Callback) {
	return func(m *tb.Callback) {
		factory := &DefaultGameFactory{repo: ef.userRepo}
		newGame := factory.NewGame(&NoobDifficulty{})
		chat := m.Message.Chat
		bot := Bot{TgBot: b, Chat: chat}
		b.Send(chat, "Get ready for the memory game!")
		showBoard(bot, chat, *newGame)
		newGame.startGame(&bot, chat)
	}
}

func (ef *endpointsFactory) Help(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		b.Send(m.Sender, HELP)
	}
}

func (ef *endpointsFactory) GlobalRating(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		users, err := ef.userRepo.GetAllUser(UserFilter{Limit: 10})
		if err != nil {
			log.Errorf("Error while getting global rating: %v", err)
			b.Send(m.Sender, "Error while getting global rating")
			return
		}
		var rating string
		rating = "Global rating TOP 10:\n"

		for i, user := range users {
			rating += fmt.Sprintf("%d. %s - %d\n", i+1, user.UserName, user.Score)
		}

		b.Send(m.Sender, rating)
	}
}

func (ef *endpointsFactory) MyRating(b *tb.Bot) func(m *tb.Message) {
	return func(m *tb.Message) {
		user, err := ef.userRepo.GetByExternalID(fmt.Sprint(m.Sender.ID))
		if err != nil {
			log.Errorf("Error while getting user rating: %v; external_id=%d", err, m.Sender.ID)
			b.Send(m.Sender, "You have no rating yet")
			return
		}

		b.Send(m.Sender, fmt.Sprintf("Your rating is %d", user.Score))
	}
}
