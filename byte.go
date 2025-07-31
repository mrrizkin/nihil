package nihil

import (
	"database/sql"
	"database/sql/driver"
)

// NilByte wraps sql.NullByte with JSON support
type NilByte sql.NullByte

// Byte creates a valid NilByte with the given value
func Byte(b byte) NilByte {
	return NilByte{Valid: true, Byte: b}
}

// ByteNil creates an invalid (null) NilByte
func ByteNil() NilByte {
	return NilByte{Valid: false}
}

// Interface implementations for nullableJSON
func (n *NilByte) isValid() bool        { return n.Valid }
func (n *NilByte) getValue() byte       { return n.Byte }
func (n *NilByte) setValid(valid bool)  { n.Valid = valid }
func (n *NilByte) setValue(value byte)  { n.Byte = value }
func (n *NilByte) scan(value any) error { return (*sql.NullByte)(n).Scan(value) }
func (n *NilByte) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Byte), nil
}

func (n *NilByte) Scan(value any) error        { return n.scan(value) }
func (n NilByte) Value() (driver.Value, error) { return n.driverValue() }

func (n NilByte) MarshalJSON() ([]byte, error)  { return marshalNullableJSON((*NilByte)(&n)) }
func (n *NilByte) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }
