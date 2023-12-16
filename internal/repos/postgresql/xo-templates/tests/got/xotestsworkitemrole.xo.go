package got

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// XoTestsWorkItemRole is the 'work_item_role' enum type from schema 'xo_tests'.
type XoTestsWorkItemRole string

// XoTestsWorkItemRole values.
const (
	// XoTestsWorkItemRolePreparer is the 'preparer' work_item_role.
	XoTestsWorkItemRolePreparer XoTestsWorkItemRole = "preparer"
	// XoTestsWorkItemRoleReviewer is the 'reviewer' work_item_role.
	XoTestsWorkItemRoleReviewer XoTestsWorkItemRole = "reviewer"
)

// Value satisfies the driver.Valuer interface.
func (xtwir XoTestsWorkItemRole) Value() (driver.Value, error) {
	return string(xtwir), nil
}

// Scan satisfies the sql.Scanner interface.
func (xtwir *XoTestsWorkItemRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*xtwir = XoTestsWorkItemRole(s)
	case string:
		*xtwir = XoTestsWorkItemRole(s)
	default:
		return fmt.Errorf("unsupported scan type for XoTestsWorkItemRole: %T", src)
	}
	return nil
}

// NullXoTestsWorkItemRole represents a null 'work_item_role' enum for schema 'xo_tests'.
type NullXoTestsWorkItemRole struct {
	XoTestsWorkItemRole XoTestsWorkItemRole
	// Valid is true if XoTestsWorkItemRole is not null.
	Valid bool
}

// Value satisfies the driver.Valuer interface.
func (nxtwir NullXoTestsWorkItemRole) Value() (driver.Value, error) {
	if !nxtwir.Valid {
		return nil, nil
	}
	return nxtwir.XoTestsWorkItemRole.Value()
}

// Scan satisfies the sql.Scanner interface.
func (nxtwir *NullXoTestsWorkItemRole) Scan(v interface{}) error {
	if v == nil {
		nxtwir.XoTestsWorkItemRole, nxtwir.Valid = "", false
		return nil
	}
	err := nxtwir.XoTestsWorkItemRole.Scan(v)
	nxtwir.Valid = err == nil
	return err
}

// ErrInvalidXoTestsWorkItemRole is the invalid XoTestsWorkItemRole error.
type ErrInvalidXoTestsWorkItemRole string

// Error satisfies the error interface.
func (err ErrInvalidXoTestsWorkItemRole) Error() string {
	return fmt.Sprintf("invalid XoTestsWorkItemRole(%s)", string(err))
}

func AllXoTestsWorkItemRoleValues() []XoTestsWorkItemRole {
	return []XoTestsWorkItemRole{
		XoTestsWorkItemRolePreparer,
		XoTestsWorkItemRoleReviewer,
	}
}
