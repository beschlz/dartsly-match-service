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

func TestPlayercountShouldNotConstruct(t *testing.T) {
	_, err := NewPlayerCount(0)

	if err == nil {
		t.Errorf("PlayerCount should not construct")
	}
}

func TestPlayercountShouldNotConstructNegative(t *testing.T) {
	_, err := NewPlayerCount(-1)

	if err == nil {
		t.Errorf("PlayerCount should not construct")
	}
}
