#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# For all commands, the working directory is the parent directory(repo root).
REPO_ROOT=$(git rev-parse --show-toplevel)
cd "${REPO_ROOT}"

GO111MODULE=on go install "k8s.io/code-generator/cmd/register-gen@v0.29.1"
GO111MODULE=on go install "k8s.io/code-generator/cmd/client-gen@v0.29.1"
GO111MODULE=on go install "k8s.io/code-generator/cmd/lister-gen@v0.29.1"
GO111MODULE=on go install "k8s.io/code-generator/cmd/informer-gen@v0.29.1"

# echo "Generating deep copy files with deepcopy-gen"
deepcopy-gen \
  -O zz_generated.deepcopy \
  --go-header-file hack/boilerplate/boilerplate.generatego.txt \
  --output-base "${REPO_ROOT}" \
  --output-package "sigs.k8s.io/about-api/" \
  --input-dirs sigs.k8s.io/about-api/api/v1alpha1 \
  --input-dirs sigs.k8s.io/about-api/api/v1beta1 \
  --trim-path-prefix "${REPO_ROOT}/sigs.k8s.io/about-api/"

echo "Generating register files with register-gen"
register-gen \
  -O zz_generated.register \
  --go-header-file hack/boilerplate/boilerplate.generatego.txt \
  --output-base "${REPO_ROOT}" \
  --output-package "sigs.k8s.io/about-api/" \
  --input-dirs sigs.k8s.io/about-api/api/v1alpha1 \
  --input-dirs sigs.k8s.io/about-api/api/v1beta1 \
  --trim-path-prefix "${REPO_ROOT}/sigs.k8s.io/about-api/"

echo "Generating versioned clientsets with client-gen"
client-gen \
  --go-header-file "hack/boilerplate/boilerplate.generatego.txt" \
  --clientset-name "versioned" \
  --input-base "sigs.k8s.io/about-api/" \
  --input "api/v1alpha1/,api/v1beta1/" \
  --output-base "${REPO_ROOT}" \
  --output-package "sigs.k8s.io/about-api/pkg/generated/clientset/" \
  --input-dirs sigs.k8s.io/about-api/api/v1alpha1,sigs.k8s.io/about-api/api/v1beta1 \
  --trim-path-prefix "${REPO_ROOT}/sigs.k8s.io/about-api/"

echo "Generating listers with lister-gen"
lister-gen \
  --go-header-file "hack/boilerplate/boilerplate.generatego.txt" \
  --output-base "${REPO_ROOT}" \
  --output-package "sigs.k8s.io/about-api/pkg/generated/listers/" \
  --input-dirs sigs.k8s.io/about-api/api/v1alpha1,sigs.k8s.io/about-api/api/v1beta1 \
  --trim-path-prefix "${REPO_ROOT}/sigs.k8s.io/about-api/"

echo "Generating informers with informer-gen"
informer-gen \
  --go-header-file "hack/boilerplate/boilerplate.generatego.txt" \
  --output-base "${REPO_ROOT}" \
  --output-package "sigs.k8s.io/about-api/pkg/generated/informers/" \
  --input-dirs sigs.k8s.io/about-api/api/v1alpha1,sigs.k8s.io/about-api/api/v1beta1 \
  --trim-path-prefix "${REPO_ROOT}/sigs.k8s.io/about-api/" \
  --versioned-clientset-package sigs.k8s.io/about-api/pkg/generated/clientset/versioned \
  --listers-package sigs.k8s.io/about-api/pkg/generated/listers
