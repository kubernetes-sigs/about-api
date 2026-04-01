#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# For all commands, the working directory is the parent directory(repo root).
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

CODEGEN_PKG=${CODEGEN_PKG:-$(go -C tools mod download -json k8s.io/code-generator | jq -r .Dir)}

. "${CODEGEN_PKG}/kube_codegen.sh"

kube::codegen::gen_helpers \
  --boilerplate hack/boilerplate/boilerplate.generatego.txt \
  .

kube::codegen::gen_register \
  --boilerplate hack/boilerplate/boilerplate.generatego.txt \
  .

kube::codegen::gen_client \
  --output-dir pkg/generated \
  --output-pkg sigs.k8s.io/about-api/pkg/generated \
  --boilerplate hack/boilerplate/boilerplate.generatego.txt \
  --with-watch \
  --with-applyconfig \
  .
