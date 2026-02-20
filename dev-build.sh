#!/usr/bin/bash

for dir in ./cmd/*/ ; do
    cd "$dir" && go build -o ../../bin/$(basename "$PWD") && cd -
done