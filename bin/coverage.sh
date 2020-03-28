#!/bin/bash

set -e

go test -coverprofile=coverage_release.part -v ./...

go test -coverprofile=coverage_panic.part -v -tags unassert_panic ./...

go test -coverprofile=coverage_stderr.part -v -tags unassert_stderr ./...

go test -coverprofile=coverage_stderr.part -v -tags unassert_test ./...

echo "mode: set" >coverage.txt
grep -h -v "mode: set" *.part >>coverage.txt
