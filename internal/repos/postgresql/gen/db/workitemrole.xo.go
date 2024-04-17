// Code generated by xo. DO NOT EDIT.

//lint:ignore

package db

import (
	"database/sql/driver"
	"fmt"
)

// WorkItemRole is the 'work_item_role' enum type from schema 'public'.
type WorkItemRole string

// WorkItemRole values.
const (
	// WorkItemRolePreparer is the 'preparer' work_item_role.
	WorkItemRolePreparer WorkItemRole = "preparer"
	// WorkItemRoleReviewer is the 'reviewer' work_item_role.
	WorkItemRoleReviewer WorkItemRole = "reviewer"
)

// Value satisfies the driver.Valuer interface.
func (wir WorkItemRole) Value() (driver.Value, error) {
	return string(wir), nil
}

// Scan satisfies the sql.Scanner interface.
func (wir *WorkItemRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*wir = WorkItemRole(s)
	case string:
		*wir = WorkItemRole(s)
	default:
		return fmt.Errorf("unsupported scan type for WorkItemRole: %T", src)
	}
	return nil
}

// ErrInvalidWorkItemRole is the invalid WorkItemRole error.
type ErrInvalidWorkItemRole string

// Error satisfies the error interface.
func (err ErrInvalidWorkItemRole) Error() string {
	return fmt.Sprintf("invalid WorkItemRole(%s)", string(err))
}

func AllWorkItemRoleValues() []WorkItemRole {
	return []WorkItemRole{
		WorkItemRolePreparer,
		WorkItemRoleReviewer,
	}
}
