package matches

import (
	"errors"

	"github.com/beschlz/dartsly-match-service/internal/core/matches/domain"
)

var InvalidMatchConfiguration = errors.New("invalid match configuration")

type MatchID string

func NewMatchID(id string) MatchID {
	return MatchID(id)
}

type MatchRepo interface {
	CreateMatch() MatchID
}

type MatchService interface {
	CreateMatch(request domain.MatchCreationRequest) (*domain.Match, error)
}
