package t_bot

import (
	"math/rand"
	"time"
)

// GameFactory interface
type GameFactory interface {
	NewGame(rows, cols int) *game
}

// DefaultGameFactory struct
type DefaultGameFactory struct {
	repo UserInfo
}

// emojis randomly selected for the game with 2 pairs 16 emojis
var emojis = []string{"🐶", "🐱", "🐭", "🐹", "🐰", "🦊", "🐻", "🐼", "🐨", "🐯"}

func (f *DefaultGameFactory) NewGame(strategy DifficultyStrategy) *game {
	// Get the board size based on the selected difficulty
	rows, cols := strategy.GetBoardSize()

	rand.Seed(time.Now().UnixNano())
	em := make([]string, 0)
	em = append(em, emojis[:(rows * cols) / 2]...)
	em = append(em, em...)
	selectedEmojis := rand.Perm(len(em))
	rand.Shuffle(len(selectedEmojis), func(i, j int) {
		selectedEmojis[i], selectedEmojis[j] = selectedEmojis[j], selectedEmojis[i]
	})

	revealed := make([][]bool, rows)
	for i := range revealed {
		revealed[i] = make([]bool, cols)
	}

	board := make([][]string, rows)
	for i := range board {
		board[i] = make([]string, cols)
		for j := range board[i] {
			board[i][j] = em[selectedEmojis[i*cols+j]]
		}
	}

	return &game{board: board, revealed: revealed, firstCol: -1, firstRow: -1, prizeScore: strategy.GetWinningScore(), repo: f.repo}
}
