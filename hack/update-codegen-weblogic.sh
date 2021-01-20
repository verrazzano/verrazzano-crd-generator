#!/bin/bash
# Copyright (c) 2020, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname $0)/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; test -d ./vendor/k8s.io/code-generator 2>/dev/null && echo ./vendor/k8s.io/code-generator || echo ../code-generator)}

GENERATED_ZZ_FILE=$SCRIPT_ROOT/pkg/apis/weblogic/v8/zz_generated.deepcopy.go
echo Remove $GENERATED_ZZ_FILE file if exist
rm -f $GENERATED_ZZ_FILE

GENERATED_CLIENT_DIR=$SCRIPT_ROOT/pkg/clientwks
echo Remove $GENERATED_CLIENT_DIR dir if exist
rm -rf $GENERATED_CLIENT_DIR

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
${CODEGEN_PKG}/generate-groups.sh "deepcopy,client,informer,lister" \
  github.com/verrazzano/verrazzano-crd-generator/pkg/clientwks github.com/verrazzano/verrazzano-crd-generator/pkg/apis \
  weblogic:v8 \
  --output-base "${GOPATH}/src" \
  --go-header-file ${SCRIPT_ROOT}/hack/custom-header.txt
