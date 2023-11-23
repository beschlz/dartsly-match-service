package matches

type Match struct {
	playerCount PlayerCount
}

func NewMatch(playerCount int) *Match {
	return &Match{
		playerCount: PlayerCount(playerCount),
	}
}
