package domain

import "testing"

func TestSetsShouldConstruct(t *testing.T) {
	validSets := []int{1, 3}

	for _, setValue := range validSets {
		_, err := NewSets(1)

		if err != nil {
			t.Errorf("Sets with a value of %d should construct", setValue)
		}
	}
}

func TestSetsShouldNotConstruct(t *testing.T) {
	validSets := []int{0, -1}

	for _, setValue := range validSets {
		_, err := NewSets(setValue)

		if err == nil {
			t.Errorf("Sets with a value of %d should not construct", setValue)
		}
	}
}
