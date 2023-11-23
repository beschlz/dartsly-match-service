package ports

import "errors"

var InvalidMatchConfiguration = errors.New("invalid match configuration")

type MatchRepo interface {
	CreateMatch()
}
