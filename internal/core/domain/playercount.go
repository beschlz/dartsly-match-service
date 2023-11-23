package matches

import "errors"

var ErrInvalidPlayerCount = errors.New("invalid player count")

type PlayerCount int

func NewPlayerCount(playerCount int) (PlayerCount, error) {
	if playerCount < 1 {
		return 0, ErrInvalidPlayerCount
	}

	return PlayerCount(playerCount), nil
}
