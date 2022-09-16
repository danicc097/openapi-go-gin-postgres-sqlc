#!/usr/bin/env bash

# TODO fix bad subst -> unexpected token
# ${"$API_PREFIX"} is not changed
existing_vars=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,)
echo "${existing_vars[*]}"
for file in $JSFOLDER; do
  envsubst $existing_vars <"$file" | sponge "$file"
done
nginx -g 'daemon off;'
