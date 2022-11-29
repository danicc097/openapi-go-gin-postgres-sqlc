//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type WorkItemRole string

const (
	WorkItemRole_Preparer WorkItemRole = "preparer"
	WorkItemRole_Reviewer WorkItemRole = "reviewer"
)

func (e *WorkItemRole) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "preparer":
		*e = WorkItemRole_Preparer
	case "reviewer":
		*e = WorkItemRole_Reviewer
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for WorkItemRole enum")
	}

	return nil
}

func (e WorkItemRole) String() string {
	return string(e)
}
