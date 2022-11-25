package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"time"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
	"github.com/swaggest/openapi-go/openapi3"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type req struct {
	ID     string `path:"id" example:"XXX-XXXXX"`
	Locale string `query:"locale" pattern:"^[a-z]{2}-[A-Z]{2}$"`
	Title  string `json:"string"`
	Amount uint   `json:"amount"`
	Items  []struct {
		Count uint   `json:"count"`
		Name  string `json:"name"`
	} `json:"items,omitempty"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
}

type resp struct {
	ID     string `json:"id" example:"XXX-XXXXX"`
	Amount uint   `json:"amount"`
	Items  []struct {
		Count uint   `json:"count"`
		Name  string `json:"name"`
	} `json:"items,omitempty"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	var structNamesList string

	flag.StringVar(&structNamesList, "struct-names", "", "comma-delimited db package struct names to generate an OpenAPI schema for")
	flag.Parse()

	structNames := strings.Split(structNamesList, ",")
	for i := range structNames {
		structNames[i] = strings.TrimSpace(structNames[i])
	}

	reflector := openapi3.Reflector{Spec: &openapi3.Spec{}}

	reflector.InterceptDefName(func(t reflect.Type, defaultDefName string) string {
		if strings.HasPrefix(defaultDefName, "Db") {
			return strings.TrimPrefix(defaultDefName, "Db")
		}

		return defaultDefName
	})

	for i, sn := range structNames {
		dummyOp := openapi3.Operation{}
		st, ok := postgen.DbStructs[sn]
		if !ok {
			log.Fatalf("struct-name %s does not exist in db package", sn)
		}
		// TODO for every field in st, set tags to be `json="<value of openapi-json tag>"`
		var fields []reflect.StructField
		for i := 0; i < reflect.TypeOf(st).NumField(); i++ {
			field := reflect.TypeOf(st).Field(i)
			if strings.HasPrefix(field.Name, "_") {
				continue
			}
			t := field.Tag.Get("openapi-json")
			if t == "" {
				log.Fatalf("field %v in struct-name %s does not have an `openapi-json` tag", field, sn)
			}
			field.Tag = reflect.StructTag(fmt.Sprintf("json=\"%s\"", t))
			fields = append(fields, field)
		}
		newSt := reflect.StructOf(fields)
		fmt.Printf("reflect.New(newSt): %#v\n", reflect.New(newSt))

		// output: {00000000-0000-0000-0000-000000000000   <nil> <nil> <nil>  <nil> [] 0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil> <nil> <nil> <nil>}
		handleError(reflector.SetJSONResponse(&dummyOp, reflect.Indirect(reflect.New(newSt)), http.StatusTeapot))
		// output: {00000000-0000-0000-0000-000000000000   <nil> <nil> <nil>  <nil> [] 0 0001-01-01 00:00:00 +0000 UTC 0001-01-01 00:00:00 +0000 UTC <nil> <nil> <nil> <nil> false false} // last booleans xo "_XXX" fields
		// handleError(reflector.SetJSONResponse(&dummyOp, st, http.StatusTeapot))
		fmt.Printf("dummyOp: %#v\n", dummyOp)
		reflector.Spec.Components.Schemas.MapOfSchemaOrRefValues[sn].Schema.MapOfAnything = map[string]interface{}{"x-db-struct": sn}
		handleError(reflector.Spec.AddOperation(http.MethodGet, "/dummy-op-"+strconv.Itoa(i), dummyOp))
		// reflector.Spec.Paths.MapOfPathItemValues["mypath"].MapOfOperationValues["method"].
	}
	s, err := reflector.Spec.MarshalYAML()
	handleError(err)

	fmt.Println(string(s))
	// os.WriteFile("openapi.test.gen.yaml", schema, 0o777)

	// fmt.Println(s.Components.Schemas.MapOfSchemaOrRefValues["Error"].Schema.Properties["code"].Schema.MapOfAnything["x-foo"])
}
