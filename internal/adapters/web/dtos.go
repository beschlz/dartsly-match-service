package web

type MatchCreationRequest struct {
	PlayerCount int `json:"player_count"`
}

type MatchCreationResponse struct {
	MatchID     string `json:"match_id"`
	PlayerCount int    `json:"player_count"`
}
