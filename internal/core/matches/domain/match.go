package domain

import "errors"

var (
	ErrInvalidMatchSetting = errors.New("invalid match setting")
)

type Match struct {
	MatchID          MatchID
	PlayerCount      PlayerCount
	CheckoutSettings CheckOutSetting
	Points           Points
	Sets             Sets
	Legs             Legs
}

type MatchCreationRequest struct {
	PlayerCount      int
	CheckOutSettings string
	Points           Points
	Sets             Sets
	Legs             Legs
}

func NewMatch(request *MatchCreationRequest) (*Match, error) {
	pCount, err := NewPlayerCount(request.PlayerCount)

	if err != nil {
		return nil, ErrInvalidMatchSetting
	}

	checkoutSettings, err := NewCheckOutSetting(request.CheckOutSettings)

	if err != nil {
		return nil, ErrInvalidMatchSetting
	}

	return &Match{
			PlayerCount:      pCount,
			CheckoutSettings: checkoutSettings,
		},
		nil
}
