package domain

import "testing"

func TestMatchShouldConstruct(t *testing.T) {
	matchCreationRequest := MatchCreationRequest{
		PlayerCount:      2,
		CheckOutSettings: "singleout",
	}

	match, err := NewMatch(&matchCreationRequest)

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if match.PlayerCount != 2 {
		t.Errorf("Expected player count to be 2, got %d", match.PlayerCount)
	}

	if string(match.CheckoutSettings) != "singleout" {
		t.Errorf("Expected checkout settings to be singleout, got %s", match.CheckoutSettings)
	}
}

func InvalidMatchShouldNotConstruct(t *testing.T) {
	matchCreationRequest := MatchCreationRequest{
		PlayerCount:      -1,
		CheckOutSettings: "singleout2",
	}

	_, err := NewMatch(&matchCreationRequest)

	if err == nil {
		t.Errorf("An invalid match should not be created")
	}
}
