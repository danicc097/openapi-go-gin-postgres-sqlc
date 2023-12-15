import (
	"database/sql/driver"
	"fmt"
)

// ExtraSchemaWorkItemRole is the 'work_item_role' enum type from schema 'extra_schema'.
type ExtraSchemaWorkItemRole string

// ExtraSchemaWorkItemRole values.
const (
	// ExtraSchemaWorkItemRolePreparer is the 'preparer' work_item_role.
	ExtraSchemaWorkItemRolePreparer ExtraSchemaWorkItemRole = "preparer"
	// ExtraSchemaWorkItemRoleReviewer is the 'reviewer' work_item_role.
	ExtraSchemaWorkItemRoleReviewer ExtraSchemaWorkItemRole = "reviewer"
)

// Value satisfies the driver.Valuer interface.
func (eswir ExtraSchemaWorkItemRole) Value() (driver.Value, error) {
	return string(eswir), nil
}

// Scan satisfies the sql.Scanner interface.
func (eswir *ExtraSchemaWorkItemRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*eswir = ExtraSchemaWorkItemRole(s)
	case string:
		*eswir = ExtraSchemaWorkItemRole(s)
	default:
		return fmt.Errorf("unsupported scan type for ExtraSchemaWorkItemRole: %T", src)
	}
	return nil
}

// NullExtraSchemaWorkItemRole represents a null 'work_item_role' enum for schema 'extra_schema'.
type NullExtraSchemaWorkItemRole struct {
	ExtraSchemaWorkItemRole ExtraSchemaWorkItemRole
	// Valid is true if ExtraSchemaWorkItemRole is not null.
	Valid bool
}

// Value satisfies the driver.Valuer interface.
func (neswir NullExtraSchemaWorkItemRole) Value() (driver.Value, error) {
	if !neswir.Valid {
		return nil, nil
	}
	return neswir.ExtraSchemaWorkItemRole.Value()
}

// Scan satisfies the sql.Scanner interface.
func (neswir *NullExtraSchemaWorkItemRole) Scan(v interface{}) error {
	if v == nil {
		neswir.ExtraSchemaWorkItemRole, neswir.Valid = "", false
		return nil
	}
	err := neswir.ExtraSchemaWorkItemRole.Scan(v)
	neswir.Valid = err == nil
	return err
}

// ErrInvalidExtraSchemaWorkItemRole is the invalid ExtraSchemaWorkItemRole error.
type ErrInvalidExtraSchemaWorkItemRole string

// Error satisfies the error interface.
func (err ErrInvalidExtraSchemaWorkItemRole) Error() string {
	return fmt.Sprintf("invalid ExtraSchemaWorkItemRole(%s)", string(err))
}

func AllExtraSchemaWorkItemRoleValues() []ExtraSchemaWorkItemRole {
	return []ExtraSchemaWorkItemRole{
		ExtraSchemaWorkItemRolePreparer,
		ExtraSchemaWorkItemRoleReviewer,
	}
}

