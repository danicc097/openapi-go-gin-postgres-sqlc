#!/bin/bash

SCHEMA_OUT="src/types/schema.d.ts"

mkdir -p src/redux/slices/gen
mkdir -p src/types

rtk-query-codegen-openapi openapi-config.json &
node generate-client-validator.js &
openapi-typescript ../openapi.yaml --output "$SCHEMA_OUT" --path-params-as-types --prettier-config .prettierrc &
wait

echo "export type schemas = components['schemas']" | cat - "$SCHEMA_OUT" >/tmp/out && mv /tmp/out "$SCHEMA_OUT"
echo "// @ts-nocheck" | cat - "$SCHEMA_OUT" >/tmp/out && mv /tmp/out "$SCHEMA_OUT"
echo "/* eslint-disable */" | cat - "$SCHEMA_OUT" >/tmp/out && mv /tmp/out "$SCHEMA_OUT"
echo "/* eslint-disable @typescript-eslint/ban-ts-comment */" | cat - "$SCHEMA_OUT" >/tmp/out && mv /tmp/out "$SCHEMA_OUT"
# TODO allow custom validate.ts baked into fork
rm src/client-validator/gen/validate.ts
find src/client-validator/gen/ -type f -exec \
  sed -i "s/from '.\/validate'/from '..\/validate'/g" {} \;
