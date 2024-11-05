package t_bot

import (
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	NewGameReplyButton = tb.ReplyButton{Text: "New Game 🎮"}
	GlobalRatingButton = tb.ReplyButton{Text: "Global Rating [top 10] 📊"}
	MyRatingButton     = tb.ReplyButton{Text: "My Rating 📈"}
	HelpButton         = tb.ReplyButton{Text: "Help ❓"}

	NoobButton   = tb.InlineButton{Text: "noob", Unique: "h1"}
	EasyButton   = tb.InlineButton{Text: "easy", Unique: "h1"}
	MediumButton = tb.InlineButton{Text: "medium", Unique: "h2"}
	HardButton   = tb.InlineButton{Text: "hard", Unique: "h3"}

	ReplyKeys = [][]tb.ReplyButton{
		{NewGameReplyButton, GlobalRatingButton},
		{MyRatingButton, HelpButton},
	}

	GameOptionReplyKeys = [][]tb.InlineButton{
		{NoobButton, EasyButton, MediumButton, HardButton},
	}
)
