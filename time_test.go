package nihil

import (
	"encoding/json"
	"testing"
	"time"
)

func TestNilTime_Constructor(t *testing.T) {
	now := time.Now()

	// Test valid value
	validTime := Time(now)
	if !validTime.Valid {
		t.Error("Expected valid time to be valid")
	}
	if !validTime.Time.Equal(now) {
		t.Errorf("Expected time value %v, got %v", now, validTime.Time)
	}

	// Test nil value
	nilTime := TimeNil()
	if nilTime.Valid {
		t.Error("Expected nil time to be invalid")
	}
}

func TestNilTime_JSONMarshaling(t *testing.T) {
	testTime := time.Date(2023, 10, 15, 14, 30, 0, 0, time.UTC)

	tests := []struct {
		name     string
		input    NilTime
		expected string
	}{
		{"valid time", Time(testTime), `"2023-10-15T14:30:00Z"`},
		{"nil time", TimeNil(), "null"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := json.Marshal(tt.input)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}

func TestNilTime_JSONUnmarshaling(t *testing.T) {
	testTime := time.Date(2023, 10, 15, 14, 30, 0, 0, time.UTC)

	tests := []struct {
		name          string
		input         string
		expectedValid bool
		expectedTime  time.Time
	}{
		{"valid time", `"2023-10-15T14:30:00Z"`, true, testTime},
		{"null", "null", false, time.Time{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result NilTime
			err := json.Unmarshal([]byte(tt.input), &result)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result.Valid != tt.expectedValid {
				t.Errorf("Expected valid=%v, got valid=%v", tt.expectedValid, result.Valid)
			}
			if tt.expectedValid && !result.Time.Equal(tt.expectedTime) {
				t.Errorf("Expected time=%v, got time=%v", tt.expectedTime, result.Time)
			}
		})
	}
}
