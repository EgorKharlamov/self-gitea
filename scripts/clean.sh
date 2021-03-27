#!/usr/bin/env bash

function main() {
    BUILD_DIR="$1"

    if [ -d "$1" ]; then
      rm -r "$BUILD_DIR"
    fi
}

main "$1"
