package nihil

import (
	"database/sql"
	"database/sql/driver"
)

// NilString wraps sql.NullString with JSON support
type NilString sql.NullString

// String creates a valid NilString with the given value
func String(s string) NilString {
	return NilString{Valid: true, String: s}
}

// StringNil creates an invalid (null) NilString
func StringNil() NilString {
	return NilString{Valid: false}
}

// Interface implementations for nullableJSON
func (n *NilString) isValid() bool         { return n.Valid }
func (n *NilString) getValue() string      { return n.String }
func (n *NilString) setValid(valid bool)   { n.Valid = valid }
func (n *NilString) setValue(value string) { n.String = value }
func (n *NilString) scan(value any) error  { return (*sql.NullString)(n).Scan(value) }
func (n *NilString) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.String, nil
}

func (n *NilString) Scan(value any) error        { return n.scan(value) }
func (n NilString) Value() (driver.Value, error) { return n.driverValue() }

func (n NilString) MarshalJSON() ([]byte, error) {
	return marshalNullableJSON((*NilString)(&n))
}
func (n *NilString) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }
