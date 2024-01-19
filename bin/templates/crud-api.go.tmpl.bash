#!/bin/bash

# shellcheck disable=SC2028,SC2154
cat <<EOF
package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gin-gonic/gin"
$(test -n "$with_project" && echo "	\"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal\"")
)

func (h *StrictHandlers) Create${pascal_name}(c *gin.Context, request Create${pascal_name}RequestObject) (Create${pascal_name}ResponseObject, error) {
	tx := GetTxFromCtx(c)

	params := request.Body.${pascal_name}CreateParams
$(test -n "$with_project" && echo "	params.ProjectID = internal.ProjectIDByName[request.ProjectName]")

	${camel_name}, err := h.svc.${pascal_name}.Create(c, tx, &params)
	if err != nil {
		renderErrorResponse(c, "Could not create ${sentence_name}", err)

		return nil, nil
	}

	res := ${pascal_name}{
		${pascal_name}: *${camel_name},
		// joins, if any
	}

	return Create${pascal_name}201JSONResponse(res), nil
}

func (h *StrictHandlers) Get${pascal_name}(c *gin.Context, request Get${pascal_name}RequestObject) (Get${pascal_name}ResponseObject, error) {
	tx := GetTxFromCtx(c)

	${camel_name}, err := h.svc.${pascal_name}.ByID(c, tx, db.${pascal_name}ID(request.${pascal_name}ID))
	if err != nil {
		renderErrorResponse(c, "Could not create ${sentence_name}", err)

		return nil, nil
	}

	res := ${pascal_name}{
		${pascal_name}: *${camel_name},
		// joins, if any
	}

	return Get${pascal_name}200JSONResponse(res), nil
}

func (h *StrictHandlers) Update${pascal_name}(c *gin.Context, request Update${pascal_name}RequestObject) (Update${pascal_name}ResponseObject, error) {
	tx := GetTxFromCtx(c)

	params := request.Body.${pascal_name}UpdateParams

	${camel_name}, err := h.svc.${pascal_name}.Update(c, tx, db.${pascal_name}ID(request.${pascal_name}ID), &params)
	if err != nil {
		renderErrorResponse(c, "Could not update ${sentence_name}", err)

		return nil, nil
	}

	res := ${pascal_name}{
		${pascal_name}: *${camel_name},
		// joins, if any
	}

	return Update${pascal_name}200JSONResponse(res), nil
}

func (h *StrictHandlers) Delete${pascal_name}(c *gin.Context, request Delete${pascal_name}RequestObject) (Delete${pascal_name}ResponseObject, error) {
	tx := GetTxFromCtx(c)

	_, err := h.svc.${pascal_name}.Delete(c, tx, db.${pascal_name}ID(request.${pascal_name}ID))
	if err != nil {
		renderErrorResponse(c, "Could not delete ${sentence_name}", err)

		return nil, nil
	}

	return Delete${pascal_name}204Response{}, nil
}
EOF
