#!/usr/bin/env bash

set -e

./clean.sh
./build.sh

rm -f wasm.zip

zip -r wasm *.wasm *.js *.html assets/**/*

butler push wasm.zip cultist-games/thats-my-spot:wasm
