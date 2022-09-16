#!/usr/bin/env bash

existing_vars=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,)
# bad subst for $_: /usr/local/bin/envsubst=Symbol.for("react.fragment... breaks js
# ideally would pass env var names to Dockerfile to only subst those
existing_vars=("${existing_vars[@]/',$_'/}")

for file in $JSFOLDER; do
  envsubst $existing_vars <"$file" | sponge "$file"
done

nginx -g 'daemon off;'
