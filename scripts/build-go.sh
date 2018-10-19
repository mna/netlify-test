#!/bin/bash

set -euo pipefail

function main {
  local func_dir="$1"
  local bin_dir="$2"

  for dir in "${func_dir}"*/; do
    local name=""; name=$(basename "${dir}")
    echo "building Go function ${name}..."
    go build -o "${bin_dir}${name}" "${dir}"
  done
}

main "$@"
