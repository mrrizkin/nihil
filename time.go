package nihil

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

// NilTime wraps sql.NullTime with JSON support
type NilTime sql.NullTime

// Time creates a valid NilTime with the given value
func Time(t time.Time) NilTime {
	return NilTime{Valid: true, Time: t}
}

// TimeNil creates an invalid (null) NilTime
func TimeNil() NilTime {
	return NilTime{Valid: false}
}

// Interface implementations for nullableJSON
func (n *NilTime) isValid() bool            { return n.Valid }
func (n *NilTime) getValue() time.Time      { return n.Time }
func (n *NilTime) setValid(valid bool)      { n.Valid = valid }
func (n *NilTime) setValue(value time.Time) { n.Time = value }
func (n *NilTime) scan(value any) error     { return (*sql.NullTime)(n).Scan(value) }
func (n *NilTime) driverValue() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Time, nil
}

func (n *NilTime) Scan(value any) error        { return n.scan(value) }
func (n NilTime) Value() (driver.Value, error) { return n.driverValue() }

func (n NilTime) MarshalJSON() ([]byte, error)  { return marshalNullableJSON((*NilTime)(&n)) }
func (n *NilTime) UnmarshalJSON(b []byte) error { return unmarshalNullableJSON(n, b) }
