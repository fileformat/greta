#!/usr/bin/env bash
#
# run tests locally
#

set -o errexit
set -o pipefail
set -o nounset

echo "INFO: starting at $(date -u +%Y-%m-%dT%H:%M:%SZ)"
go test -cover -coverprofile=coverage.txt -timeout 30s -run "^TestGreta$" .

go tool cover -html=coverage.txt -o coverage.html

echo "INFO: complete at $(date -u +%Y-%m-%dT%H:%M:%SZ)"