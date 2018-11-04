package mgr

import (
	"testing"
)

func BenchmarkGenerateOrderID(b *testing.B) {
	var currentID = autoID
	var result int64
	for i := 0; i < b.N; i++ {
		result = generateOrderID()
	}
	expected := currentID + int64(b.N)
	if result != expected {
		b.Errorf("unexpected result: got %v want %v", result, expected)
	}
}

func TestValidateLatLong(t *testing.T) {
	tt := []struct {
		latitude, longitude string
		shouldPass          bool
	}{
		{"+90.0", "-127.554334", true},
		{"45", "180", true},
		{"-90", "-180", true},
		{"-90.000", "-180.0000", true},
		{"+90", "+180", true},
		{"47.1231231", "179.99999999", true},
		{"-90.", "-180.", false},
		{"+90.1", "-100.111", false},
		{"-91", "123.456", false},
		{"045", "180", false},
		{"heap", "", false},
		{"", "", false},
	}
	for _, tc := range tt {
		result := validateLatLong(tc.latitude, tc.longitude)
		expected := tc.shouldPass
		if result != expected {
			t.Errorf("%s, %s unexpected result: got %v want %v", tc.latitude, tc.longitude, result, expected)
		}
	}
}

func TestCalculateDistance(t *testing.T) {
	tt := struct {
		start, end []string
		expected   int
	}{
		[]string{"", ""},
		[]string{"", ""},
		1,
	}
	result := calculateDistance(tt.start, tt.end)
	expected := tt.expected
	if result != expected {
		t.Errorf("unexpected result: got %v want %v", result, expected)
	}
}
