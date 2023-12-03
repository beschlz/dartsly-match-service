package domain

type Legs int

func NewLegs(legs int) (Legs, error) {
	if legs < 1 {
		return 0, ErrInvalidMatchSetting
	}

	return Legs(legs), nil
}
