package persistance

import "github.com/beschlz/dartsly-match-service/internal/core/matches"

type MatchRepo struct {
}

func NewMatchRepo() matches.MatchRepo {
	return &MatchRepo{}
}

func (m *MatchRepo) CreateMatch() {

}
