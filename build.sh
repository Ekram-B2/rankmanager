#!/bin/bash

set -e -x
go get -t -d -v ./...
go build -o build/rank-manager .