#!/usr/bin/env bash

set -e
./clean.sh
env GOOS=js GOARCH=wasm go build -o ./thatsmyspot.wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .