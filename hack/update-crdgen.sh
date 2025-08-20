#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo "Generating with controller-gen"
GO111MODULE=on go install "sigs.k8s.io/controller-tools/cmd/controller-gen@v0.17.2"

# Unify the crds used by helm chart and the installation scripts
controller-gen crd paths=./pkg/apis/... output:crd:dir=./config/crd/
