#!/usr/bin/env bash

set -e

./clean.sh
./build.sh

butler push ./build/wasm cultist-games/thats-my-spot:wasm

butler push ./build/windows cultist-games/thats-my-spot:windows

butler push ./build/linux cultist-games/thats-my-spot:linux
