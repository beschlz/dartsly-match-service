package domain

import "testing"

func TestPointsShouldConstruct(t *testing.T) {
	Points, err := NewPoints(301)

	if err != nil || Points != 301 {
		t.Errorf("Points should construct")
	}

	Points, err = NewPoints(501)

	if err != nil || Points != 501 {
		t.Errorf("Points should construct")
	}
}

func TestPointsShouldNotConstruct(t *testing.T) {
	_, err := NewPoints(302)

	if err == nil {
		t.Errorf("Points should not construct")
	}
}
