#!/usr/bin/env bash
#
# run tests locally
#

set -o errexit
set -o pipefail
set -o nounset

echo "INFO: starting at $(date -u +%Y-%m-%dT%H:%M:%SZ)"

if [ -f "./greta" ]; then
    echo "INFO: removing old build of greta"
    rm ./greta
fi
COMMIT=local
#$(git rev-parse HEAD --short)
LASTMOD=$(date -u +%Y-%m-%dT%H:%M:%SZ)
VERSION=local
BUILTBY=run.sh

echo "INFO: building new greta"
go build \
    -ldflags "-X main.commit=$COMMIT -X main.date=$LASTMOD -X main.version=$VERSION -X main.builtBy=$BUILTBY -extldflags '-static'" \
    -o ./greta \
    ./greta.go

if [ ! -f "./greta" ]; then
    echo "ERROR: failed to build greta"
    exit 1
fi

echo "INFO: running $(greta --version)"

echo "INFO: running greta with arguments: $@"
./greta "$@"

echo "INFO: complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"