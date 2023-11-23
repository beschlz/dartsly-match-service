package domain

import (
	"testing"
)

func TestPlayercountShouldConstruct(t *testing.T) {
	PlayerCount, err := NewPlayerCount(1)

	if err != nil || PlayerCount != 1 {
		t.Errorf("PlayerCount should construct")
	}
}
