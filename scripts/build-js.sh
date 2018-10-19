#!/bin/bash

set -euo pipefail

function main {
  local func_dir="$1"
  local bin_dir="$2"
	cp --verbose "${func_dir}"*.js "${bin_dir}"
}

main "$@"
