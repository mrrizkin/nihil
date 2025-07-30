package nihil

import (
	"encoding/json"
	"testing"
)

func TestNilBool_Constructor(t *testing.T) {
	// Test valid true
	validTrue := Bool(true)
	if !validTrue.Valid {
		t.Error("Expected valid bool to be valid")
	}
	if !validTrue.Bool {
		t.Error("Expected bool value to be true")
	}

	// Test valid false
	validFalse := Bool(false)
	if !validFalse.Valid {
		t.Error("Expected valid bool to be valid")
	}
	if validFalse.Bool {
		t.Error("Expected bool value to be false")
	}

	// Test nil value
	nilBool := BoolNil()
	if nilBool.Valid {
		t.Error("Expected nil bool to be invalid")
	}
}

func TestNilBool_JSONMarshaling(t *testing.T) {
	tests := []struct {
		name     string
		input    NilBool
		expected string
	}{
		{"valid true", Bool(true), "true"},
		{"valid false", Bool(false), "false"},
		{"nil bool", BoolNil(), "null"},
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

func TestNilBool_JSONUnmarshaling(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValid bool
		expectedValue bool
	}{
		{"true", "true", true, true},
		{"false", "false", true, false},
		{"null", "null", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result NilBool
			err := json.Unmarshal([]byte(tt.input), &result)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result.Valid != tt.expectedValid {
				t.Errorf("Expected valid=%v, got valid=%v", tt.expectedValid, result.Valid)
			}
			if result.Bool != tt.expectedValue {
				t.Errorf("Expected value=%v, got value=%v", tt.expectedValue, result.Bool)
			}
		})
	}
}
