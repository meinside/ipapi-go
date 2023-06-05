package ipapi

import (
	"testing"
)

const (
	test_ip      = "8.8.8.8"
	test_country = "US"
)

func TestGetLocation(t *testing.T) {
	result, err := GetLocation(test_ip)

	if err != nil {
		t.Errorf("Failed to get location: %s", err)
	} else {
		if result.Country != test_country {
			t.Errorf("Expected fetched country to be `%s`, got %s", test_country, result.Country)
		}
	}
}
