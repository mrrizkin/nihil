package nihil

import (
	"encoding/json"
	"testing"
)

func TestNilString_Constructor(t *testing.T) {
	// Test valid value
	validString := String("hello")
	if !validString.Valid {
		t.Error("Expected valid string to be valid")
	}
	if validString.String != "hello" {
		t.Errorf("Expected string value 'hello', got '%s'", validString.String)
	}

	// Test nil value
	nilString := StringNil()
	if nilString.Valid {
		t.Error("Expected nil string to be invalid")
	}
}

func TestNilString_JSONMarshaling(t *testing.T) {
	tests := []struct {
		name     string
		input    NilString
		expected string
	}{
		{"regular string", String("hello"), `"hello"`},
		{"empty string", String(""), `""`},
		{"string with quotes", String(`say "hello"`), `"say \"hello\""`},
		{"nil string", StringNil(), "null"},
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

func TestNilString_JSONUnmarshaling(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValid bool
		expectedValue string
	}{
		{"regular string", `"hello"`, true, "hello"},
		{"empty string", `""`, true, ""},
		{"null", "null", false, ""},
		{"string with escapes", `"hello\nworld"`, true, "hello\nworld"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result NilString
			err := json.Unmarshal([]byte(tt.input), &result)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result.Valid != tt.expectedValid {
				t.Errorf("Expected valid=%v, got valid=%v", tt.expectedValid, result.Valid)
			}
			if result.String != tt.expectedValue {
				t.Errorf("Expected value='%s', got value='%s'", tt.expectedValue, result.String)
			}
		})
	}
}
