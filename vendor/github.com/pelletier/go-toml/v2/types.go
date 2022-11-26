package toml

import (
	"encoding"
	"reflect"
	"time"
)

var (
	timeType               = reflect.TypeOf(time.Time{})
	textMarshalerType      = reflect.TypeOf(new(encoding.TextMarshaler)).Elem()
	textUnmarshalerType    = reflect.TypeOf(new(encoding.TextUnmarshaler)).Elem()
	mapStringInterfaceType = reflect.TypeOf(map[string]interface{}{})
	sliceInterfaceType     = reflect.TypeOf([]interface{}{})
	stringType             = reflect.TypeOf("")
)
