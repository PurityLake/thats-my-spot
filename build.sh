#!/usr/bin/env bash

set -e
./clean.sh

mkdir -p ./build/{wasm,linux,windows}/

# WASM
env GOOS=js GOARCH=wasm go build -o ./thatsmyspot.wasm
cp ./thatsmyspot.wasm ./build/wasm
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
cp wasm_exec.js ./build/wasm
cp -r ./assets ./build/wasm
cp ./index.html ./build/wasm

# Linux
env GOOS=linux go build -o ./build/linux/thatsmyspot
cp -r ./assets ./build/linux

# Windows
env GOOS=windows go build -o ./build/windows/thatsmyspot.exe
cp -r ./assets ./build/windows

