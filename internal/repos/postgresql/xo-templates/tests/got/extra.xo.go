package got

// Code generated by xo. DO NOT EDIT.

import (
	"fmt"
)

func newPointer[T any](v T) *T {
	return &v
}

type XoError struct {
	Entity string
	Err    error
}

// Error satisfies the error interface.
func (e *XoError) Error() string {
	return fmt.Sprintf("%s %v", e.Entity, e.Err)
}

// Unwrap satisfies the unwrap interface.
func (err *XoError) Unwrap() error {
	return err.Err
}

type TableEntity string

type Filter struct {
	// Typ is the field type. It is one of: string, number, integer, boolean, date-time
	// Arrays and objects are ignored for default filter generation
	Typ string `json:"type"`
	// Db is the corresponding db column name
	Db       string `json:"db"`
	Nullable bool   `json:"nullable"`
}
