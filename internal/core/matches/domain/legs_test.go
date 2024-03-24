package domain

import "testing"

func TestLegsShouldConstruct(t *testing.T) {
	validLegs := []int{1, 3}

	for _, legValue := range validLegs {
		_, err := NewLegs(1)

		if err != nil {
			t.Errorf("Legs with a value of %d should construct", legValue)
		}
	}
}

func TestLegsShouldNotConstruct(t *testing.T) {
	invalidLegs := []int{0, -1}

	for _, legValue := range invalidLegs {
		_, err := NewLegs(legValue)

		if err == nil {
			t.Errorf("Legs with a value of %d should not construct", legValue)
		}
	}
}
