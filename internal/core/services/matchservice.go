package matches

import (
	matches "github.com/beschlz/dartsly-match-service/internal/core/ports"
	"github.com/beschlz/dartsly-match-service/internal/ports"
)

type matchService struct {
	matchRepo ports.MatchRepo
}

func (m *matchService) CreateMatch() string {
	return "match created"
}

func NewMatchService(matchRepo ports.MatchRepo) matches.MatchService {
	return &matchService{
		matchRepo: matchRepo,
	}
}
