#! /bin/bash
set -e

dep ensure -v

for pkg in `go list ./...`; do
    go test -v -race -coverprofile=profile.out $pkg
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
