#!/usr/bin/env bash

set -e

./clean.sh
go run github.com/hajimehoshi/wasmserve@latest .
