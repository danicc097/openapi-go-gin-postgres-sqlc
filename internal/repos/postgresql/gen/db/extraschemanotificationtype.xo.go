import (
	"database/sql/driver"
	"fmt"
)

// ExtraSchemaNotificationType is the 'notification_type' enum type from schema 'extra_schema'.
type ExtraSchemaNotificationType string

// ExtraSchemaNotificationType values.
const (
	// ExtraSchemaNotificationTypePersonal is the 'personal' notification_type.
	ExtraSchemaNotificationTypePersonal ExtraSchemaNotificationType = "personal"
	// ExtraSchemaNotificationTypeGlobal is the 'global' notification_type.
	ExtraSchemaNotificationTypeGlobal ExtraSchemaNotificationType = "global"
)

// Value satisfies the driver.Valuer interface.
func (esnt ExtraSchemaNotificationType) Value() (driver.Value, error) {
	return string(esnt), nil
}

// Scan satisfies the sql.Scanner interface.
func (esnt *ExtraSchemaNotificationType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*esnt = ExtraSchemaNotificationType(s)
	case string:
		*esnt = ExtraSchemaNotificationType(s)
	default:
		return fmt.Errorf("unsupported scan type for ExtraSchemaNotificationType: %T", src)
	}
	return nil
}

// NullExtraSchemaNotificationType represents a null 'notification_type' enum for schema 'extra_schema'.
type NullExtraSchemaNotificationType struct {
	ExtraSchemaNotificationType ExtraSchemaNotificationType
	// Valid is true if ExtraSchemaNotificationType is not null.
	Valid bool
}

// Value satisfies the driver.Valuer interface.
func (nesnt NullExtraSchemaNotificationType) Value() (driver.Value, error) {
	if !nesnt.Valid {
		return nil, nil
	}
	return nesnt.ExtraSchemaNotificationType.Value()
}

// Scan satisfies the sql.Scanner interface.
func (nesnt *NullExtraSchemaNotificationType) Scan(v interface{}) error {
	if v == nil {
		nesnt.ExtraSchemaNotificationType, nesnt.Valid = "", false
		return nil
	}
	err := nesnt.ExtraSchemaNotificationType.Scan(v)
	nesnt.Valid = err == nil
	return err
}

// ErrInvalidExtraSchemaNotificationType is the invalid ExtraSchemaNotificationType error.
type ErrInvalidExtraSchemaNotificationType string

// Error satisfies the error interface.
func (err ErrInvalidExtraSchemaNotificationType) Error() string {
	return fmt.Sprintf("invalid ExtraSchemaNotificationType(%s)", string(err))
}

func AllExtraSchemaNotificationTypeValues() []ExtraSchemaNotificationType {
	return []ExtraSchemaNotificationType{
		ExtraSchemaNotificationTypePersonal,
		ExtraSchemaNotificationTypeGlobal,
	}
}

