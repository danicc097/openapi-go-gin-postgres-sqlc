package main

import (
	"fmt"
	"go/importer"
	"go/token"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"
	"unsafe"

	// kinopenapi3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/swaggest/openapi-go/openapi3"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//go:linkname typelinks reflect.typelinks
func typelinks() (sections []unsafe.Pointer, offset [][]int32)

//go:linkname add reflect.add
func add(p unsafe.Pointer, x uintptr, whySafe string) unsafe.Pointer

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

// clear && go build -o gen-schema cmd/gen-schema/main.go && ./gen-schema -env .env.dev
func main() {
	reflector := openapi3.Reflector{Spec: &openapi3.Spec{}}

	// we could load the existing spec to reflector.Spec: https://pkg.go.dev/github.com/swaggest/openapi-go/openapi3#example-Spec.UnmarshalYAML
	// and if "x-db-struct" found in an OPERATION.
	// NOTE: comments are gone. should print result to openapi.gen.yaml and use that since this is in gen/postgen step
	schemaBlob, err := os.ReadFile("openapi.yaml")
	if err != nil {
		log.Fatalf("openapi spec: %s", err)
	}
	if err := reflector.Spec.UnmarshalYAML(schemaBlob); err != nil {
		log.Fatalf("Spec.UnmarshalYAML: %s", err)
	}

	// we can edit an existing op by getting all operations with "x-db-struct", trace back to the openapi3.Operation
	putOp := openapi3.Operation{}

	// FIXME this wont work, we will have access only to things referenced by User. the rest of structs are missing
	// fmt.Println(db.User{}) // dummy call to dynamically get types in package
	// modelsPkg := "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	var dbTypes []reflect.Type
	var dbTypesdd []string
	sections, offsets := typelinks()
	for i, base := range sections {
		for _, offset := range offsets[i] {
			typeAddr := add(base, uintptr(offset), "")
			typ := reflect.TypeOf(*(*interface{})(unsafe.Pointer(&typeAddr)))
			// if typ.Kind() != reflect.Struct {
			// 	continue
			// }
			dbTypes = append(dbTypes, typ)
			dbTypesdd = append(dbTypesdd, typ.String()+"|"+typ.Kind().String())
			// fmt.Println("--------------")
			// fmt.Println(typ.PkgPath())
			// fmt.Println(typ.String())
			// reflect.New()
		}
	}
	dbFiles, err := os.ReadDir("./internal/repos/postgresql/gen/db")
	if err != nil {
		log.Fatal(err)
	}

	fs := token.NewFileSet()
	for _, f := range dbFiles {
		if f.IsDir() {
			continue
		}
		i, err := f.Info()
		if err != nil {
			log.Fatalf("info")
		}

		fs.AddFile(f.Name(), fs.Base(), int(i.Size()))
	}

	defaultImporter := importer.ForCompiler(fs, "source", nil)
	pkg, err := defaultImporter.Import("github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db")
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}
	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
	// fix missing structs! Activity, etc. for some reason dont appear
	// for i, base := range sections {
	// 	for _, offset := range offsets[i] {
	// 		typeAddr := add(base, uintptr(offset), "")
	// 		typ := reflect.TypeOf(*(*interface{})(unsafe.Pointer(&typeAddr)))
	// 		if typ.Kind() != reflect.Pointer {
	// 			continue
	// 		}
	// 		if typ.Elem().Kind() != reflect.Struct {
	// 			continue
	// 		}
	// 		st := typ.Elem()
	// 		dbTypes = append(dbTypes, st)
	// 		dbTypesdd = append(dbTypesdd, st.String()+"|"+st.Kind().String())
	// 		// fmt.Println("--------------")
	// 		fmt.Println(st.PkgPath())
	// 		// fmt.Println(st.String())
	// 		// reflect.New()
	// 	}
	// }
	// fmt.Println(dbTypesdd)

	handleError(reflector.SetRequest(&putOp, new(req), http.MethodPut))
	// handleError(reflector.SetJSONResponse(&putOp, new(db.User), http.StatusOK))
	// handleError(reflector.SetJSONResponse(&putOp, new([]db.User), http.StatusConflict))
	handleError(reflector.Spec.AddOperation(http.MethodPut, "/things/{id}", putOp))
	// reflector.Spec.Paths.MapOfPathItemValues["fefse"].MapOfOperationValues["test"].
	schema, err := reflector.Spec.MarshalYAML()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(schema))

	os.WriteFile("openapi.test.gen.yaml", schema, 0o777)
	fmt.Println("saved spec")

	// var s openapi3.Spec

	// TODO pgx types instead of null.* package. For time use pgtype.Date /pgtype.Timestamptz etc.
	// pgtype supports json (un)marshalling like `null.v4`
	// openapi-go needs some kind of annotation to tell its nullable: true and format: date-time or whatever
	// https://github.com/jackc/pgtype
	// have to merge with our own (done easily but comments are lost. output to openapi.gen.yaml, we are in postgen step so doesnt matter)
	// schemaBlob, err := os.ReadFile("openapi.yaml")
	// if err != nil {
	// 	log.Fatalf("error opening schema file: %s", err)
	// }

	// oas, err := rest.ReadOpenAPI("openapi.yaml")
	// if err != nil {
	// 	log.Fatalf("ReadOpenAPI: %s", err)
	// }
	// // oas.Components.SecuritySchemes = kinopenapi3.SecuritySchemes{} // error

	// // fmt.Println(string(schemaBlob))
	// specWithoutSec, err := oas.MarshalJSON()
	// if err != nil {
	// 	log.Fatalf("oas.MarshalJSON: %s", err)
	// }
	// // fmt.Println(string(specWithoutSec))

	// if err := s.UnmarshalYAML(specWithoutSec); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(s.Info.Title)
	// fmt.Println(s.Info.Title)
	// fmt.Println(s.Components.Schemas.MapOfSchemaOrRefValues["Error"].Schema.Properties["code"].Schema.MapOfAnything["x-foo"])
}
