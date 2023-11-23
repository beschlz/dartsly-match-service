package domain

import "errors"

var (
	InvalidMatchSetting = errors.New("invalid match setting")
)

type CheckOutSetting string

var CheckOut = struct {
	SingleOut CheckOutSetting
	DoubleOut CheckOutSetting
}{
	SingleOut: "singleout",
	DoubleOut: "doubleout",
}

func NewCheckOutSetting(setting CheckOutSetting) (CheckOutSetting, error) {
	if setting != CheckOut.SingleOut && setting != CheckOut.DoubleOut {

		return "", InvalidMatchSetting
	}

	return CheckOutSetting(setting), nil
}

type Match struct {
	PlayerCount      PlayerCount
	CheckoutSettings CheckOutSetting
}

type MatchCreationRequest struct {
	PlayerCount      int
	CheckOutSettings CheckOutSetting
}

func NewMatch(request *MatchCreationRequest) (*Match, error) {
	pCount, err := NewPlayerCount(request.PlayerCount)

	if err != nil {
		return nil, InvalidMatchSetting
	}

	checkoutSettings, err := NewCheckOutSetting(request.CheckOutSettings)

	if err != nil {
		return nil, InvalidMatchSetting
	}

	return &Match{
			PlayerCount:      pCount,
			CheckoutSettings: checkoutSettings,
		},
		nil
}
