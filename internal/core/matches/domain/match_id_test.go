package domain

import "testing"

func TestMatchIdShouldConstruct(t *testing.T) {
	id := "test"
	matchId := NewMatchId(id)

	if string(matchId) != id {
		t.Errorf("expected match id to be %s, but was %s", id, matchId)
	}
}
