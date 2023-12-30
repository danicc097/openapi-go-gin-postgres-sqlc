# shellcheck disable=SC2028,SC2154
echo "
x-require-authenticated: &x-require-authenticated
x-error-response: &x-error-response
paths:
$(if test -n "$with_project"; then
  echo "  /project/{projectName}/${kebab_name}/:"
else
  echo "  /${kebab_name}/:"
fi)
    post:
      summary: create ${sentence_name}.
      !!merge <<: *x-require-authenticated
      operationId: Create${pascal_name}
      x-required-scopes:
        - ${kebab_name}:create
$(if test -n "$with_project"; then
  echo "
      parameters:
        - \$ref: '#/components/parameters/ProjectName'"
fi)
      requestBody:
        content:
          application/json:
            schema:
              \$ref: '#/components/schemas/Create${pascal_name}Request'
        required: true
      responses:
        201:
          description: Success.
          content:
            application/json:
              schema:
                \$ref: '#/components/schemas/${pascal_name}'
        !!merge <<: *x-error-response
      tags:
        - ${camel_name}
  /${kebab_name}/{id}:
    get:
      summary: get ${sentence_name}.
      !!merge <<: *x-require-authenticated
      operationId: Get${pascal_name}
      parameters:
        - \$ref: '#/components/parameters/SerialID'
      responses:
        200:
          description: Success.
          content:
            application/json:
              schema:
                \$ref: '#/components/schemas/${pascal_name}'
        !!merge <<: *x-error-response
      tags:
        - ${camel_name}
    patch:
      summary: update ${sentence_name}.
      !!merge <<: *x-require-authenticated
      operationId: Update${pascal_name}
      x-required-scopes:
        - ${kebab_name}:edit
      parameters:
        - \$ref: '#/components/parameters/SerialID'
      requestBody:
        content:
          application/json:
            schema:
              \$ref: '#/components/schemas/Update${pascal_name}Request'
        required: true
      responses:
        200:
          description: Success.
          content:
            application/json:
              schema:
                \$ref: '#/components/schemas/${pascal_name}'
        !!merge <<: *x-error-response
      tags:
        - ${camel_name}
    delete:
      summary: delete $name.
      !!merge <<: *x-require-authenticated
      operationId: Delete${pascal_name}
      x-required-scopes:
        - ${kebab_name}:delete
      parameters:
        - \$ref: '#/components/parameters/SerialID'
      responses:
        204:
          description: Success.
        !!merge <<: *x-error-response
      tags:
        - ${camel_name}
components:
  schemas:
    Create${pascal_name}Request:
      x-postgen-struct: RestCreate${pascal_name}Request
    Update${pascal_name}Request:
      x-postgen-struct: RestUpdate${pascal_name}Request
    ${pascal_name}:
      x-postgen-struct: Rest${pascal_name}
"
