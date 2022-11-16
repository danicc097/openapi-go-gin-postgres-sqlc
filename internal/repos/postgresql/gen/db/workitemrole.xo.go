package db

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql/driver"
	"fmt"
)

// WorkItemRole is the 'work_item_role' enum type from schema 'public'.
type WorkItemRole uint16

// WorkItemRole values.
const (
	// WorkItemRolePreparer is the 'preparer' work_item_role.
	WorkItemRolePreparer WorkItemRole = 1
	// WorkItemRoleReviewer is the 'reviewer' work_item_role.
	WorkItemRoleReviewer WorkItemRole = 2
)

// String satisfies the fmt.Stringer interface.
func (wir WorkItemRole) String() string {
	switch wir {
	case WorkItemRolePreparer:
		return "preparer"
	case WorkItemRoleReviewer:
		return "reviewer"
	}
	return fmt.Sprintf("WorkItemRole(%d)", wir)
}

// MarshalText marshals WorkItemRole into text.
func (wir WorkItemRole) MarshalText() ([]byte, error) {
	return []byte(wir.String()), nil
}

// UnmarshalText unmarshals WorkItemRole from text.
func (wir *WorkItemRole) UnmarshalText(buf []byte) error {
	switch str := string(buf); str {
	case "preparer":
		*wir = WorkItemRolePreparer
	case "reviewer":
		*wir = WorkItemRoleReviewer
	default:
		return ErrInvalidWorkItemRole(str)
	}
	return nil
}

// Value satisfies the driver.Valuer interface.
func (wir WorkItemRole) Value() (driver.Value, error) {
	return wir.String(), nil
}

// Scan satisfies the sql.Scanner interface.
func (wir *WorkItemRole) Scan(v interface{}) error {
	switch buf := v.(type) {
	case []byte:
		return wir.UnmarshalText(buf)
	case string:
		return wir.UnmarshalText([]byte(buf))
	}

	return ErrInvalidWorkItemRole(fmt.Sprintf("%T", v))
}

// NullWorkItemRole represents a null 'work_item_role' enum for schema 'public'.
type NullWorkItemRole struct {
	WorkItemRole WorkItemRole
	// Valid is true if WorkItemRole is not null.
	Valid bool
}

// Value satisfies the driver.Valuer interface.
func (nwir NullWorkItemRole) Value() (driver.Value, error) {
	if !nwir.Valid {
		return nil, nil
	}
	return nwir.WorkItemRole.Value()
}

// Scan satisfies the sql.Scanner interface.
func (nwir *NullWorkItemRole) Scan(v interface{}) error {
	if v == nil {
		nwir.WorkItemRole, nwir.Valid = 0, false
		return nil
	}
	err := nwir.WorkItemRole.Scan(v)
	nwir.Valid = err == nil
	return err
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
