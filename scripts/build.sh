#!/usr/bin/env bash

function main() {
    BUILD_DIR="$1"
    APP_NAME="$2"

    checkIfNotExists "$BUILD_DIR"
    renameBuildIfExist "$BUILD_DIR" "$APP_NAME"

    go build -o "$BUILD_DIR"/"$APP_NAME" cmd/app/*.go
}

function checkIfNotExists() {
  if [ ! -d "$1" ]; then
    mkdir -p "$1"
  fi
}

function renameBuildIfExist() {
  if [ -f "$1/$2" ]; then
    mv "$1/$2" "$1/$2_$(date "+%Y%m%d_%H%M%S")"
  fi
}

main "$1" "$2"
