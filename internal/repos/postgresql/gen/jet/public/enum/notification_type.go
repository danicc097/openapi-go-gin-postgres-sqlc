//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package enum

import "github.com/go-jet/jet/v2/postgres"

var NotificationType = &struct {
	Personal postgres.StringExpression
	Global   postgres.StringExpression
}{
	Personal: postgres.NewEnumValue("personal"),
	Global:   postgres.NewEnumValue("global"),
}
