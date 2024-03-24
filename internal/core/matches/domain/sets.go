package domain

type Sets int

func NewSets(sets int) (Sets, error) {
	if sets < 1 {
		return 0, ErrInvalidMatchSetting
	}

	return Sets(sets), nil
}
