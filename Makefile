# Copyright (c) 2020, Oracle Corporation and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

GO ?= GO111MODULE=on go
CODEGEN_PATH = k8s.io/code-generator
CRD_DIR = deploy/crds

.PHONY: go-build
go-build: go-mod
	./hack/update-codegen.sh
	./hack/update-codegen-weblogic.sh
	./hack/update-codegen-istio.sh
	./hack/update-codegen-coherence.sh
	operator-sdk generate k8s
	operator-sdk generate crds

	# These crds are generated but not needed
	rm deploy/crds/coherence.oracle.com*
	rm deploy/crds/networking.istio.io*
	rm deploy/crds/weblogic.oracle*

	# Add copyright headers to the operator-sdk
	# generated CRDs
	./hack/add-crd-header.sh

.PHONY: go-mod
go-mod:
	$(GO) mod vendor

	# go mod vendor only copies the .go files.  Also need
	# to populate the k8s.io/code-generator folder with the
	# scripts for generating informer/lister code

	# Obtain k8s.io/code-generator version
	$(eval codeGenVer=$(shell grep "code-generator =>" go.mod | awk '{print $$4}'))

	# Add the required files into the vendor folder
	cp ${GOPATH}/pkg/mod/${CODEGEN_PATH}@${codeGenVer}/generate-groups.sh vendor/${CODEGEN_PATH}/generate-groups.sh
	chmod +x vendor/${CODEGEN_PATH}/generate-groups.sh
	cp -R ${GOPATH}/pkg/mod/${CODEGEN_PATH}@${codeGenVer}/cmd/defaulter-gen vendor/${CODEGEN_PATH}/cmd/defaulter-gen
	chmod -R +w vendor/${CODEGEN_PATH}/cmd/defaulter-gen


