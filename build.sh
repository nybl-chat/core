#!/bin/bash

export CGO_ENABLED=0
GOOS=linux

for dir in ./cmd/*/; do
    name=$(basename "$dir")
    go build -o "./bin/$name" "./$dir"
done