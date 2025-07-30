package nihil

import (
	"database/sql"
	"database/sql/driver"
)

// Numeric types - using type aliases for cleaner code
type (
	NilFloat64 sql.NullFloat64
	NilInt16   sql.NullInt16
	NilInt32   sql.NullInt32
	NilInt64   sql.NullInt64
)

// Float64 constructors
func Float64(f float64) NilFloat64 { return NilFloat64{Valid: true, Float64: f} }
func Float64Nil() NilFloat64       { return NilFloat64{Valid: false} }

// Int16 constructors
func Int16(i int16) NilInt16 { return NilInt16{Valid: true, Int16: i} }
func Int16Nil() NilInt16     { return NilInt16{Valid: false} }

// Int32 constructors
func Int32(i int32) NilInt32 { return NilInt32{Valid: true, Int32: i} }
func Int32Nil() NilInt32     { return NilInt32{Valid: false} }

// Int64 constructors
func Int64(i int64) NilInt64 { return NilInt64{Valid: true, Int64: i} }
func Int64Nil() NilInt64     { return NilInt64{Valid: false} }

// NilFloat64 implementations
func (n *NilFloat64) isValid() bool          { return n.Valid }
func (n *NilFloat64) getValue() float64      { return n.Float64 }
func (n *NilFloat64) setValid(valid bool)    { n.Valid = valid }
func (n *NilFloat64) setValue(value float64) { n.Float64 = value }
func (n *NilFloat64) scan(value any) error   { return (*sql.NullFloat64)(n).Scan(value) }
func (n *NilFloat64) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Float64, nil
}

func (n *NilFloat64) Scan(value any) error        { return n.scan(value) }
func (n NilFloat64) Value() (driver.Value, error) { return n.driverValue() }
func (n NilFloat64) MarshalJSON() ([]byte, error) {
	return marshalNullableJSON((*NilFloat64)(&n))
}
func (n *NilFloat64) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }

// NilInt16 implementations
func (n *NilInt16) isValid() bool        { return n.Valid }
func (n *NilInt16) getValue() int16      { return n.Int16 }
func (n *NilInt16) setValid(valid bool)  { n.Valid = valid }
func (n *NilInt16) setValue(value int16) { n.Int16 = value }
func (n *NilInt16) scan(value any) error { return (*sql.NullInt16)(n).Scan(value) }
func (n *NilInt16) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int16, nil
}

func (n *NilInt16) Scan(value any) error        { return n.scan(value) }
func (n NilInt16) Value() (driver.Value, error) { return n.driverValue() }

func (n NilInt16) MarshalJSON() ([]byte, error)  { return marshalNullableJSON((*NilInt16)(&n)) }
func (n *NilInt16) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }

// NilInt32 implementations
func (n *NilInt32) isValid() bool        { return n.Valid }
func (n *NilInt32) getValue() int32      { return n.Int32 }
func (n *NilInt32) setValid(valid bool)  { n.Valid = valid }
func (n *NilInt32) setValue(value int32) { n.Int32 = value }
func (n *NilInt32) scan(value any) error { return (*sql.NullInt32)(n).Scan(value) }
func (n *NilInt32) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int32, nil
}

func (n *NilInt32) Scan(value any) error        { return n.scan(value) }
func (n NilInt32) Value() (driver.Value, error) { return n.driverValue() }

func (n NilInt32) MarshalJSON() ([]byte, error)  { return marshalNullableJSON((*NilInt32)(&n)) }
func (n *NilInt32) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }

// NilInt64 implementations
func (n *NilInt64) isValid() bool        { return n.Valid }
func (n *NilInt64) getValue() int64      { return n.Int64 }
func (n *NilInt64) setValid(valid bool)  { n.Valid = valid }
func (n *NilInt64) setValue(value int64) { n.Int64 = value }
func (n *NilInt64) scan(value any) error { return (*sql.NullInt64)(n).Scan(value) }
func (n *NilInt64) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}

func (n *NilInt64) Scan(value any) error        { return n.scan(value) }
func (n NilInt64) Value() (driver.Value, error) { return n.driverValue() }

func (n NilInt64) MarshalJSON() ([]byte, error)  { return marshalNullableJSON((*NilInt64)(&n)) }
func (n *NilInt64) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }
