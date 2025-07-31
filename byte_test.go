package nihil

import (
	"encoding/json"
	"testing"
)

func TestNilByte_Constructor(t *testing.T) {
	// Test valid value
	validByte := Byte(42)
	if !validByte.Valid {
		t.Error("Expected valid byte to be valid")
	}
	if validByte.Byte != 42 {
		t.Errorf("Expected byte value 42, got %d", validByte.Byte)
	}

	// Test nil value
	nilByte := ByteNil()
	if nilByte.Valid {
		t.Error("Expected nil byte to be invalid")
	}
}

func TestNilByte_JSONMarshaling(t *testing.T) {
	tests := []struct {
		name     string
		input    NilByte
		expected string
	}{
		{"valid byte", Byte(255), "255"},
		{"nil byte", ByteNil(), "null"},
		{"zero byte", Byte(0), "0"},
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

func TestNilByte_JSONUnmarshaling(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedValid bool
		expectedValue byte
	}{
		{"valid byte", "42", true, 42},
		{"null", "null", false, 0},
		{"zero", "0", true, 0},
		{"max byte", "255", true, 255},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result NilByte
			err := json.Unmarshal([]byte(tt.input), &result)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result.Valid != tt.expectedValid {
				t.Errorf("Expected valid=%v, got valid=%v", tt.expectedValid, result.Valid)
			}
			if result.Byte != tt.expectedValue {
				t.Errorf("Expected value=%d, got value=%d", tt.expectedValue, result.Byte)
			}
		})
	}
}

func TestNilByte_DriverValue(t *testing.T) {
	// Test valid value
	validByte := Byte(123)
	val, err := validByte.Value()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != int64(123) {
		t.Errorf("Expected 123, got %v", val)
	}
	// Test nil value
	nilByte := ByteNil()
	val, err = nilByte.Value()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != nil {
		t.Errorf("Expected nil, got %v", val)
	}
}
