#!/bin/bash
 
set -e -u -x
 
export GOPATH=$PWD/gopath
export PATH=$PWD/gopath/bin:$PATH

echo "Running tests..."
go test -v ./...
