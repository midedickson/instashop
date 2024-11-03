package constants

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

// Money represents a monetary value stored in cents/pennies
type Money int64

// Value implements the driver.Valuer interface for database storage
func (m Money) Value() (driver.Value, error) {
	return int64(m), nil
}

// Scan implements the sql.Scanner interface for database retrieval
func (m *Money) Scan(value interface{}) error {
	if value == nil {
		*m = 0
		return nil
	}

	switch v := value.(type) {
	case int64:
		*m = Money(v)
	case float64:
		*m = Money(v * 100)
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return err
		}
		*m = Money(f * 100)
	default:
		return fmt.Errorf("cannot scan type %T into Money", value)
	}
	return nil
}

// String returns the money value as a formatted string
func (m Money) String() string {
	return fmt.Sprintf("%.2f", float64(m)/100)
}
