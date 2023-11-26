package web

import "github.com/beschlz/dartsly-match-service/internal/core/matches/domain"

func mapToDomainObject(request MatchCreationRequest) domain.MatchCreationRequest {
	return domain.MatchCreationRequest{
		PlayerCount:      request.PlayerCount,
		CheckOutSettings: "singleout",
	}
}

func toDto(createdMatch *domain.Match) MatchCreationResponse {

	return MatchCreationResponse{
		MatchID:     string(createdMatch.MatchID),
		PlayerCount: int(createdMatch.PlayerCount),
	}
}
