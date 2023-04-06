/*
gen-schema generates OpenAPI v3 schema portions from code.
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/swaggest/openapi-go/openapi3"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type (
	St1 struct {
		Id int `json:"ID"`
	}
	St2 struct {
		Id int `json:"ID"`
	}
	St3 struct {
		Id int `json:"ID"`
	}
	St4 struct {
		Id int `json:"ID"`
	}
	St5 struct {
		Id int `json:"ID"`
	}
	St6 struct {
		Id int `json:"ID"`
	}
	St7 struct {
		Id int `json:"ID"`
	}
	St8 struct {
		Id int `json:"ID"`
	}
	St9 struct {
		Id int `json:"ID"`
	}
	St10 struct {
		Id int `json:"ID"`
	}
)

var PublicStructs = map[string]any{
	"St1":  St1{},
	"St2":  St2{},
	"St3":  St3{},
	"St4":  St4{},
	"St5":  St5{},
	"St6":  St6{},
	"St7":  St7{},
	"St8":  St8{},
	"St9":  St9{},
	"St10": St10{},
}

func main() {
	structNames := []string{
		"St1",
		"St2",
		"St3",
		"St4",
		"St5",
		"St6",
		"St7",
		"St8",
		"St9",
		"St10",
	}

	reflector := openapi3.Reflector{Spec: &openapi3.Spec{}}

	reflector.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
		return defaultDefName
	})

	for i, sn := range structNames {
		dummyOp := openapi3.Operation{}
		st, ok := PublicStructs[sn]
		if !ok {
			log.Fatalf("struct-name %s does not exist in PublicStructs", sn)
		}
		fmt.Printf("sn: %v\n", sn)
		handleError(reflector.SetJSONResponse(&dummyOp, st, http.StatusTeapot))

		handleError(reflector.Spec.AddOperation(http.MethodGet, "/dummy-op-"+strconv.Itoa(i), dummyOp))

		// printCurrentSpec(reflector)
		printCurrentSchemas(reflector)

		reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[sn].Schema.MapOfAnything = map[string]interface{}{"x-postgen-struct": sn}
	}
}

func printCurrentSchemas(reflector openapi3.Reflector) {
	schemas := make([]string, 0, len(reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues))
	for k := range reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues {
		schemas = append(schemas, k)
	}
	fmt.Printf("schemas (len=%d): %v\n", len(schemas), schemas)
}

func printCurrentSpec(reflector openapi3.Reflector) {
	currentSpec, _ := reflector.Spec.MarshalYAML()
	fmt.Printf("currentSpec:\n%v\n", string(currentSpec))
}
