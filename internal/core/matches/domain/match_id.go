package domain

type MatchID string

func NewMatchId(id string) MatchID {
	return MatchID(id)
}
