package format

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/fatih/structs"
)

// PrintJSONByTag marshals to JSON.
func PrintJSON(obj any) {
	bytes, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Println(string(bytes))
}

// PrintJSONByTag marshals to JSON by a specified tag.
func PrintJSONByTag(obj any, tag string) {
	switch reflect.TypeOf(obj).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(obj)
		for i := 0; i < s.Len(); i++ {
			PrintJSONByTag(s.Index(i).Interface(), tag)
		}
	default:
		printJSONByTagSingle(obj, tag)
	}
}

func printJSONByTagSingle(obj any, tag string) {
	var s *structs.Struct
	if reflect.TypeOf(obj).Kind() == reflect.Ptr {
		s = structs.New(reflect.ValueOf(obj).Elem().Interface())
	} else {
		s = structs.New(obj)
	}
	s.TagName = tag
	m := s.Map()
	bytes, _ := json.MarshalIndent(m, "  ", "  ")
	fmt.Println(string(bytes))
}
