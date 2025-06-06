#!/usr/bin/env sh

set -e

OUT_DIR="${PWD}/pkg"

rm -rf "${OUT_DIR}"
mkdir -p "${OUT_DIR}"

docker run --rm --user "$(id -u)" --volume "${PWD}":"${PWD}" --workdir "${PWD}" tool-protoc:latest \
  --proto_path "${PWD}" \
  --go_out "${OUT_DIR}" \
  --go_opt paths=source_relative \
  --go-grpc_out "${OUT_DIR}" \
  --go-grpc_opt paths=source_relative \
  spec/v1/service.proto
