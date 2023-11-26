package web

type MatchCreationRequest struct {
	PlayerCount  int    `json:"player_count"`
	CheckoutMode string `json:"checkout_mode"`
	Points       int    `json:"points"`
}

type MatchCreationResponse struct {
	MatchID      string `json:"match_id"`
	PlayerCount  int    `json:"player_count"`
	CheckoutMode string `json:"checkout_mode"`
	Points       int    `json:"points"`
}
