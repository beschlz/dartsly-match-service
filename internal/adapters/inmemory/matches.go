package inmemory

import (
	"github.com/beschlz/dartsly-match-service/internal/core/matches"
	"github.com/google/uuid"
)

type dbMatch struct {
}

type MatchRepo struct {
	matches map[string]dbMatch
}

func NewMatchRepo() matches.MatchRepo {
	repo := &MatchRepo{}
	repo.matches = make(map[string]dbMatch)

	return repo
}

func (m *MatchRepo) CreateMatch() matches.MatchID {
	id := uuid.New().String()
	m.matches[id] = dbMatch{}

	return matches.NewMatchID(id)
}
