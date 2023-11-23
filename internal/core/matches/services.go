package matches

import "github.com/beschlz/dartsly-match-service/internal/core/matches/domain"

type matchService struct {
	matchRepo MatchRepo
}

func (m *matchService) CreateMatch(request domain.MatchCreationRequest) (domain.Match, error) {
	match, err := domain.NewMatch(&request)

	return *match, err
}

func NewMatchService(matchRepo MatchRepo) MatchService {
	return &matchService{
		matchRepo: matchRepo,
	}
}
