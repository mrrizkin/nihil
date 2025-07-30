package nihil

import (
	"encoding/json"
	"testing"
)

func TestNilInt32_Constructor(t *testing.T) {
	// Test valid value
	validInt := Int32(42)
	if !validInt.Valid {
		t.Error("Expected valid int32 to be valid")
	}
	if validInt.Int32 != 42 {
		t.Errorf("Expected int32 value 42, got %d", validInt.Int32)
	}

	// Test nil value
	nilInt := Int32Nil()
	if nilInt.Valid {
		t.Error("Expected nil int32 to be invalid")
	}
}

func TestNilInt32_JSONMarshaling(t *testing.T) {
	tests := []struct {
		name     string
		input    NilInt32
		expected string
	}{
		{"positive int", Int32(123), "123"},
		{"negative int", Int32(-456), "-456"},
		{"zero", Int32(0), "0"},
		{"nil int", Int32Nil(), "null"},
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

func TestNilFloat64_Constructor(t *testing.T) {
	// Test valid value
	validFloat := Float64(3.14)
	if !validFloat.Valid {
		t.Error("Expected valid float64 to be valid")
	}
	if validFloat.Float64 != 3.14 {
		t.Errorf("Expected float64 value 3.14, got %f", validFloat.Float64)
	}

	// Test nil value
	nilFloat := Float64Nil()
	if nilFloat.Valid {
		t.Error("Expected nil float64 to be invalid")
	}
}

func TestNilFloat64_JSONMarshaling(t *testing.T) {
	tests := []struct {
		name     string
		input    NilFloat64
		expected string
	}{
		{"positive float", Float64(123.45), "123.45"},
		{"negative float", Float64(-67.89), "-67.89"},
		{"zero", Float64(0.0), "0"},
		{"nil float", Float64Nil(), "null"},
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

func TestAllNumericConstructors(t *testing.T) {
	// Test Int16
	int16Val := Int16(16)
	if !int16Val.Valid || int16Val.Int16 != 16 {
		t.Error("Int16 constructor failed")
	}

	int16Nil := Int16Nil()
	if int16Nil.Valid {
		t.Error("Int16Nil constructor failed")
	}

	// Test Int64
	int64Val := Int64(64)
	if !int64Val.Valid || int64Val.Int64 != 64 {
		t.Error("Int64 constructor failed")
	}

	int64Nil := Int64Nil()
	if int64Nil.Valid {
		t.Error("Int64Nil constructor failed")
	}
}
