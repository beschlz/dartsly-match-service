package persistance

import "github.com/beschlz/dartsly-match-service/internal/ports"

type MatchRepo struct {
}

func NewMatchRepo() ports.MatchRepo {
	return &MatchRepo{}
}

func (m *MatchRepo) CreateMatch() {

}
