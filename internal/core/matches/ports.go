package matches

import (
	"errors"

	"github.com/beschlz/dartsly-match-service/internal/core/matches/domain"
)

var InvalidMatchConfiguration = errors.New("invalid match configuration")

type MatchRepo interface {
	CreateMatch()
}

type MatchService interface {
	CreateMatch(request domain.MatchCreationRequest) (domain.Match, error)
}
