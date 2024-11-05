package t_bot

import (
	"fmt"
	"gopkg.in/tucnak/telebot.v2"
)

// Game struct
type game struct {
	board      [][]string
	revealed   [][]bool
	firstRow   int
	firstCol   int
	observers  []Observer
	repo       UserInfo
	prizeScore int64
}

func showBoard(bot Bot, chat *telebot.Chat, g game) {
	var message string
	for i, row := range g.board {
		for j, emoji := range row {
			if g.revealed[i][j] {
				message += emoji + " "
			} else {
				message += "❓ "
			}
		}
		message += "\n"
	}

	// Send the new message
	bot.TgBot.Send(chat, message)
}

func (g *game) allUnrevealed() {
	for i := range g.revealed {
		for j := range g.revealed[i] {
			g.revealed[i][j] = false
		}
	}
}

func (g *game) startGame(bot *Bot, chat *telebot.Chat) {
	g.AddObserver(bot) // Add the bot as an observer

	bot.TgBot.Handle(telebot.OnText, func(m *telebot.Message) {
		if g.firstRow == -1 {
			// First selection
			row, col := parseSelection(m.Text)
			if isValidSelection(g.board, row, col) {
				g.firstRow = row
				g.firstCol = col
				g.revealed[row][col] = true
				showBoard(*bot, chat, *g)
				bot.TgBot.Send(chat, "Remember this emoji!")
			} else {
				bot.TgBot.Send(chat, "Invalid selection. Try again (format: row col)")
			}
		} else {
			// Second selection
			row, col := parseSelection(m.Text)
			if isValidSelection(g.board, row, col) && g.board[g.firstRow][g.firstCol] == g.board[row][col] {
				// Match
				g.revealed[row][col] = true
				showBoard(*bot, chat, *g)
				g.NotifyObservers("Match!")
				if isGameOver(*g) {
					g.NotifyObservers(fmt.Sprintf("Game over! You won %d points!", g.prizeScore))
					updateUserScore(g.repo, *chat, g.prizeScore)
					return
				}
				g.firstRow = -1
				g.firstCol = -1
			} else {
				// No match
				g.revealed[row][col] = true
				showBoard(*bot, chat, *g)
				g.allUnrevealed()
				showBoard(*bot, chat, *g)
				g.revealed[row][col] = false
				g.firstRow = -1
				g.firstCol = -1
				g.NotifyObservers("No match. Try again!")
			}
		}
	})
}

func parseSelection(text string) (int, int) {
	var row, col int
	fmt.Sscanf(text, "%d %d", &row, &col)
	return row - 1, col - 1
}

func isValidSelection(board [][]string, row, col int) bool {
	return row >= 0 && row < len(board) && col >= 0 && col < len(board[0])
}

func isGameOver(g game) bool {
	for i, row := range g.board {
		for j, _ := range row {
			if !g.revealed[i][j] {
				return false
			}
		}
	}
	return true
}

func updateUserScore(repo UserInfo, chat telebot.Chat, score int64) {
	user, err := repo.GetByExternalID(fmt.Sprintf("%d", chat.ID))
	if err != nil {
		user = &User{
			ExternalID: fmt.Sprintf("%d", chat.ID),
			FirstName:  chat.FirstName,
			LastName:   chat.LastName,
			UserName:   chat.Username,
			Score:      score,
		}
		repo.CreateUser(user)
	} else {
		user.Score += score
		repo.UpdateUser(user.ID, user)
	}
}
