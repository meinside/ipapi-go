package ipapi

import (
	"testing"
)

func TestGetLocation(t *testing.T) {
	result, err := GetLocation("8.8.8.8")

	if err != nil {
		t.Errorf("Failed to get location: %s", err)
	} else {
		if result.Country != "US" {
			t.Errorf("Expected fetched country to be `US`, got %s", result.Country)
		}
	}
}
