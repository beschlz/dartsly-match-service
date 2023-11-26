package domain

type Points int

func NewPoints(points int) (Points, error) {
	if points != 301 && points != 501 {
		return 0, ErrInvalidMatchSetting
	}

	return Points(points), nil
}
