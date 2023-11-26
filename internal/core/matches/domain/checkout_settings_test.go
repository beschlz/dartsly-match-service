package domain

import "testing"

func TestCheckoutSettingsShouldConstrut(t *testing.T) {
	list := []string{"singleout", "doubleout"}

	for _, setting := range list {
		checkoutSettings, err := NewCheckOutSetting(setting)

		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}

		if string(checkoutSettings) != setting {
			t.Errorf("Expected checkoutSettings to be %s, got %s", setting, checkoutSettings)
		}
	}
}

func TestCheckoutSettingsShouldNotConstruct(t *testing.T) {
	_, err := NewCheckOutSetting("invalid")

	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
