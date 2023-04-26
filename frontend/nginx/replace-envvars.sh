#!/usr/bin/env bash

# Accepts a list of envvars to substitute: "$ENV1[,$ENV2]"
# defaulting to substituting the complete env

existing_vars="$1"
if [[ -z $existing_vars ]]; then
  existing_vars=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,)
  # bad subst for $_: /usr/local/bin/envsubst=Symbol.for("react.fragment... breaks js
  # ideally would pass env var names to Dockerfile to only subst those
  existing_vars="${existing_vars/',$_'/}"
fi

for file in $ENV_REPLACE_GLOB; do
  envsubst "$existing_vars" <"$file" | sponge "$file"
done
