package matches

import (
	"testing"

	"github.com/beschlz/dartsly-match-service/internal/core/matches/domain"
)

type MockedMatchRepo struct {
	MatchRepo
}

func (m *MockedMatchRepo) CreateMatch() MatchID {
	return "testid"
}

func TestCreatedMatchIsTheSameAsSpecified(t *testing.T) {
	service := NewMatchService(&MockedMatchRepo{})

	matchRequest := domain.MatchCreationRequest{
		PlayerCount:      2,
		CheckOutSettings: "singleout",
	}

	createdMatch, err := service.CreateMatch(matchRequest)

	if err != nil {
		t.Errorf("Error while creating match: %s", err)
	}

	if int(createdMatch.PlayerCount) != matchRequest.PlayerCount || string(createdMatch.CheckoutSettings) != matchRequest.CheckOutSettings {
		t.Errorf("PlayerCount is not the same as specified: %d", createdMatch.PlayerCount)
	}
}

func TestMatchShouldNotBeCreate(t *testing.T) {
	service := NewMatchService(&MockedMatchRepo{})

	invalidMatchSettings := domain.MatchCreationRequest{
		PlayerCount:      -1,
		CheckOutSettings: "singleout2",
	}

	_, err := service.CreateMatch(invalidMatchSettings)

	if err == nil {
		t.Errorf("Invalid Match Specifications must not lead to a created match")
	}
}
