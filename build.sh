#!/bin/bash

export CGO_ENABLED=0

go test ./...

go install ./...
