package matches

import "testing"

func TestMatchIdShouldConstruct(t *testing.T) {
	id := NewMatchID("123")

	if id != "123" {
		t.Errorf("Expected id to be 123, got %s", id)
	}
}
