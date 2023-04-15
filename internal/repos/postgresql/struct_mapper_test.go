package postgresql

import (
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/stretchr/testify/assert"
)

type InnerStruct struct {
	InnerField1 int
	InnerField2 string
}

type StStruct struct {
	Field1 *string
	Field2 int
	Field3 bool
	Field4 InnerStruct
}

type UpdateStStruct struct {
	Field1 **string
	Field2 *int
	Field3 *bool
}

func TestUpdateEntityWithParams(t *testing.T) {
	st := StStruct{
		Field1: pointers.New("st"),
		Field2: 42,
		Field3: true,
		Field4: InnerStruct{
			InnerField1: 10,
			InnerField2: "st",
		},
	}

	stCopy := st

	stUpdateParams := UpdateStStruct{
		Field1: pointers.New(pointers.New("new string")),
		Field3: pointers.New(false),
	}

	updateEntityWithParams(&stCopy, &stUpdateParams)

	assert.Equal(t, stCopy.Field1, *stUpdateParams.Field1)
	assert.Equal(t, stCopy.Field3, *stUpdateParams.Field3)

	// ensure that the original is not modified for some reason
	assert.Equal(t, stCopy.Field2, st.Field2)
	assert.Equal(t, stCopy.Field4, st.Field4)
}
