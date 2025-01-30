package game

const (
	difficultyEasyMaxCriteriaId     uint8 = 17
	difficultyStandardMaxCriteriaId uint8 = 22
)

const (
	EasyDifficulty     Difficulty = 0
	StandardDifficulty Difficulty = 1
	HardDifficulty     Difficulty = 2
)

// The difficulty of a game
type Difficulty uint8
