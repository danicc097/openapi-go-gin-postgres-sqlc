package db

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// ExtraSchemaWorkItemRole is the 'work_item_role' enum type from schema 'extra_schema'.
type ExtraSchemaWorkItemRole string

// ExtraSchemaWorkItemRole values.
const (
	// ExtraSchemaWorkItemRoleExtraPreparer is the 'extra_preparer' work_item_role.
	ExtraSchemaWorkItemRoleExtraPreparer ExtraSchemaWorkItemRole = "extra_preparer"
	// ExtraSchemaWorkItemRoleExtraReviewer is the 'extra_reviewer' work_item_role.
	ExtraSchemaWorkItemRoleExtraReviewer ExtraSchemaWorkItemRole = "extra_reviewer"
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

// ErrInvalidExtraSchemaWorkItemRole is the invalid ExtraSchemaWorkItemRole error.
type ErrInvalidExtraSchemaWorkItemRole string

// Error satisfies the error interface.
func (err ErrInvalidExtraSchemaWorkItemRole) Error() string {
	return fmt.Sprintf("invalid ExtraSchemaWorkItemRole(%s)", string(err))
}

func AllExtraSchemaWorkItemRoleValues() []ExtraSchemaWorkItemRole {
	return []ExtraSchemaWorkItemRole{
		ExtraSchemaWorkItemRoleExtraPreparer,
		ExtraSchemaWorkItemRoleExtraReviewer,
	}
}
