#!/bin/bash
set -euo pipefail

if [ -z "${MUX_TOKEN_ID:-}" ]
then
      echo "MUX_TOKEN_ID not set"
      exit 255
fi

if [ -z "${MUX_TOKEN_SECRET:-}" ]
then
      echo "MUX_TOKEN_SECRET not set"
      exit 255
fi

# go run ./examples/video/latest/latest.go

VIDEO_TESTS=./examples/video/*/exercise*.go
for f in $VIDEO_TESTS
do
  echo "========== Running $f =========="
    go run $f
done

# DATA_TESTS=./examples/data/*/exercise*.go
# for f in $DATA_TESTS
# do
#   echo "========== Running $f =========="
#     go run $f
# done
