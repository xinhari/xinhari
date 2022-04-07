#!/bin/bash
# set -x

failed=0
go mod edit -replace xinhari.com/xinhari=xinhari.com/xinhari@$1 
# basic test, build the binary
go build
failed=$?
if [ $failed -gt 0 ]; then
    exit $failed
fi
# unit tests
IN_TRAVIS_CI=yes go test -v ./...
# TODO integration tests