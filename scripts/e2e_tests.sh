#!/bin/bash

go clean --testcache

go build -o ./e2e .

go test -v ./e2e...
if [[ $? -ne 0 ]]; then
    echo "e2e tests failed" >&2
    exit 1
fi