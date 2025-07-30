package nihil

import (
	"database/sql/driver"
	"encoding/json"
)

// nullableJSON is a generic interface for types that can handle null JSON values
// This interface is used internally to provide consistent behavior across all nullable types
type nullableJSON[T any] interface {
	isValid() bool
	getValue() T
	setValid(bool)
	setValue(T)
	scan(value any) error
	driverValue() (driver.Value, error)
}

// marshalNullableJSON is a generic helper for JSON marshaling
// It handles the common pattern of marshaling either the value or null
func marshalNullableJSON[T any](n nullableJSON[T]) ([]byte, error) {
	if n.isValid() {
		return json.Marshal(n.getValue())
	}
	return json.Marshal(nil)
}

// unmarshalNullableJSON is a generic helper for JSON unmarshaling
// It handles the common pattern of checking for null and unmarshaling the value
func unmarshalNullableJSON[T any](n nullableJSON[T], b []byte) error {
	if string(b) == "null" {
		n.setValid(false)
		return nil
	}

	var value T
	if err := json.Unmarshal(b, &value); err != nil {
		return err
	}

	n.setValue(value)
	n.setValid(true)
	return nil
}
