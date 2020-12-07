#!/bin/bash
# Copyright (c) 2020, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

# Set a default value of TCP to protocol, as a workaround for https://github.com/kubernetes-sigs/controller-tools/issues/529
set -o errexit
set -o pipefail

CRD_DIR=$(dirname $0)/../deploy/crds
sed -i -r 's/^( +)  .*or SCTP\. Defaults to "TCP"\./\0\n\1default: TCP/' $CRD_DIR/verrazzano.io_verrazzanomodels_crd.yaml
sed -i -r 's/^( +)  "TCP"\./\0\n\1default: TCP/' $CRD_DIR/verrazzano.io_verrazzanomodels_crd.yaml
