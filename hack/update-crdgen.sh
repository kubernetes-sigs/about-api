#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

echo "Generating with controller-gen"

# Unify the crds used by helm chart and the installation scripts
"$(go -C tools tool -n sigs.k8s.io/controller-tools/cmd/controller-gen)" crd paths=./api/... output:crd:dir=./config/crd/
