#!/usr/bin/env bash
set -euo pipefail
# Alternative for invoking using native Go with zero install and benefiting from Go caching.
export MAGEFILE_CACHE="${PWD}/.cache/magefile"
export GOFLAGS=""

if [[ ! -f ".cache/bin/mage" ]]; then
    go build -o .cache/bin/mage ./mage.go
fi

exec .cache/bin/mage -v "$@"
