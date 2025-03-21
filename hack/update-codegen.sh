#!/bin/bash
# Copyright 2024 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

readonly SCRIPT_ROOT=$(cd $(dirname ${BASH_SOURCE})/.. && pwd)
echo "SCRIPT_ROOT ${SCRIPT_ROOT}"
cd ${SCRIPT_ROOT}

readonly GOFLAGS="-mod=mod"
readonly GOPATH="$(mktemp -d)"

export GOFLAGS GOPATH

# Even when modules are enabled, the code-generator tools always write to
# a traditional GOPATH directory, so fake on up to point to the current
# workspace.
mkdir -p "$GOPATH/src/github.com/GoogleCloudPlatform"
ln -s "${SCRIPT_ROOT}" "$GOPATH/src/github.com/GoogleCloudPlatform/gke-networking-api"

echo "${SCRIPT_ROOT} ${GOPATH}/src/github.com/GoogleCloudPlatform/gke-networking-api"
cd ${SCRIPT_ROOT}

# https://raw.githubusercontent.com/kubernetes/code-generator/release-1.30/kube_codegen.sh
source "${SCRIPT_ROOT}/hack/kube_codegen.sh"

THIS_PKG="github.com/GoogleCloudPlatform/gke-networking-api"

for crd in "network" "gcpfirewall" "nodetopology"; do
  echo "Generating $crd CRD client"
  kube::codegen::gen_client \
      --with-watch \
      --output-dir "${SCRIPT_ROOT}/client/$crd" \
      --output-pkg "${THIS_PKG}/client/$crd" \
      --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
      --one-input-api "$crd" \
      "${SCRIPT_ROOT}/apis"

  echo "Generating $crd CRD artifacts"
  go run sigs.k8s.io/controller-tools/cmd/controller-gen crd \
    object:headerFile="${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    paths="${SCRIPT_ROOT}/apis/$crd/..." \
    output:crd:dir="${SCRIPT_ROOT}/config/crds"
done


# kube_codegen does not currently support register-gen.
# As a workaround, we invoke the register-gen script directly.
(
  # To support running this script from anywhere, first cd into this directory,
  # and then install with forced module mode on and fully qualified name.
  cd "$(dirname "${0}")"
  GO111MODULE=on go install k8s.io/code-generator/cmd/register-gen
)

# Go installs the above commands to get installed in $GOBIN if defined, and $GOPATH/bin otherwise:
GOBIN="$(go env GOBIN)"
gobin="${GOBIN:-$(go env GOPATH)/bin}"

for crd_with_version in "network/v1" "gcpfirewall/v1" "nodetopology/v1"; do
  echo "Generating register for CRD $crd_with_version"
  "${gobin}/register-gen" \
      "${SCRIPT_ROOT}/apis/$crd_with_version" \
      --go-header-file "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
      --output-file zz_generated.register.go
done

# controller-gen doesn't currently generate YAML boilerplate.
# As a workaround, we concatenate this boilerplate using cat.
for file in "${SCRIPT_ROOT}/config/crds/"/*; do
  cat "${SCRIPT_ROOT}/hack/boilerplate.yaml.txt" "$file" > temp && mv temp "$file"
done
