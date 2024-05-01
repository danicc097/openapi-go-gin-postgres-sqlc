
// Code generated by xo. DO NOT EDIT.

//lint:ignore

package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	"github.com/lib/pq"
	"github.com/lib/pq/hstore"

	"github.com/google/uuid"

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

