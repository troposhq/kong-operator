#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..

$(dirname ${BASH_SOURCE})/generate-groups.sh "deepcopy,client,informer,lister" \
  kong-operator/pkg/client kong-operator/pkg/apis \
  kong:v1alpha1 \
  --output-base "$(dirname ${BASH_SOURCE})/../.." \
  --go-header-file ${SCRIPT_ROOT}/hack/boilerplate.go.txt
