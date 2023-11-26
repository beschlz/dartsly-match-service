package domain

type CheckOutSetting string

const (
	SingleOut CheckOutSetting = "singleout"
	DoubleOut CheckOutSetting = "doubleout"
)

func NewCheckOutSetting(setting string) (CheckOutSetting, error) {
	if setting != string(SingleOut) && setting != string(DoubleOut) {

		return "", ErrInvalidMatchSetting
	}

	return CheckOutSetting(setting), nil
}
