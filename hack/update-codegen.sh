#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# For all commands, the working directory is the parent directory(repo root).
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

GO111MODULE=on go install "sigs.k8s.io/controller-tools/cmd/controller-gen@v0.16.5"
GO111MODULE=on go install "k8s.io/code-generator/cmd/register-gen@v0.29.1"
GO111MODULE=on go install "k8s.io/code-generator/cmd/client-gen@v0.29.1"
GO111MODULE=on go install "k8s.io/code-generator/cmd/lister-gen@v0.29.1"
GO111MODULE=on go install "k8s.io/code-generator/cmd/informer-gen@v0.29.1"

# set GOPATH environment variables for commands(register-gen,client-gen,lister-gen,informer-gen).
# Otherwise the generated files will be output to './' if $GOPATH is not set.
export GOPATH=$(go env GOPATH | awk -F ':' '{print $1}')

echo "Generating deep copy files with controller-gen"
controller-gen object:headerFile="hack/boilerplate/boilerplate.generatego.txt" paths="./api/v1alpha1/..."

echo "Generating with register-gen"
register-gen \
  --go-header-file hack/boilerplate/boilerplate.generatego.txt \
  --input-dirs=api/v1alpha1 \
  --output-package=api/v1alpha1 \
  --output-file-base=zz_generated.register

echo "Generating with client-gen"
client-gen \
  --go-header-file hack/boilerplate/boilerplate.generatego.txt \
  --input-base="" \
  --input=api/v1alpha1 \
  --output-package=generated/clientset \
  --clientset-name=versioned

echo "Generating with lister-gen"
lister-gen \
  --go-header-file hack/boilerplate/boilerplate.generatego.txt \
  --input-dirs=api/v1alpha1 \
  --output-package=generated/listers

echo "Generating with informer-gen"
informer-gen \
  --go-header-file hack/boilerplate/boilerplate.generatego.txt \
  --input-dirs=api/v1alpha1 \
  --versioned-clientset-package=generated/clientset/versioned \
  --listers-package=generated/listers \
  --output-package=generated/informers
