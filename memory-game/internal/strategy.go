package t_bot

type DifficultyStrategy interface {
	GetBoardSize() (int, int)
	GetWinningScore() int64
}

type EasyDifficulty struct{}

func (e *EasyDifficulty) GetBoardSize() (int, int) {
	return 3, 4
}

func (e *EasyDifficulty) GetWinningScore() int64 {
	return 100
}
type MediumDifficulty struct{}

func (m *MediumDifficulty) GetBoardSize() (int, int) {
	return 4, 4
}

func (m *MediumDifficulty) GetWinningScore() int64 {
	return 200
}

type HardDifficulty struct{}

func (h *HardDifficulty) GetBoardSize() (int, int) {
	return 5, 4
}

func (h *HardDifficulty) GetWinningScore() int64 {
	return 300
}

type NoobDifficulty struct{}

func (h *NoobDifficulty) GetBoardSize() (int, int) {
	return 2, 2
}

func (h *NoobDifficulty) GetWinningScore() int64 {
	return 1
}
