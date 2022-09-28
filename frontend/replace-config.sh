#!/usr/bin/env bash

existing_vars=$(printenv | awk -F= '{print $1}' | sed 's/^/\$/g' | paste -sd,)
existing_vars=("${existing_vars[@]/',$_'/}")
echo "existing_vars are ${existing_vars[*]}"
echo "new config:"
envsubst $existing_vars <config.template.json
envsubst $existing_vars <config.template.json >config.json
