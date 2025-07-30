package nihil

import (
	"database/sql"
	"database/sql/driver"
)

// NilBool wraps sql.NullBool with JSON support
type NilBool sql.NullBool

// Bool creates a valid NilBool with the given value
func Bool(b bool) NilBool {
	return NilBool{Valid: true, Bool: b}
}

// BoolNil creates an invalid (null) NilBool
func BoolNil() NilBool {
	return NilBool{Valid: false}
}

// Interface implementations for nullableJSON
func (n *NilBool) isValid() bool        { return n.Valid }
func (n *NilBool) getValue() bool       { return n.Bool }
func (n *NilBool) setValid(valid bool)  { n.Valid = valid }
func (n *NilBool) setValue(value bool)  { n.Bool = value }
func (n *NilBool) scan(value any) error { return (*sql.NullBool)(n).Scan(value) }
func (n *NilBool) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Bool, nil
}

func (n *NilBool) Scan(value any) error        { return n.scan(value) }
func (n NilBool) Value() (driver.Value, error) { return n.driverValue() }

func (n NilBool) MarshalJSON() ([]byte, error)  { return marshalNullableJSON((*NilBool)(&n)) }
func (n *NilBool) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }
