# Copyright (c) 2020, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

GO ?= GO111MODULE=on go
CODEGEN_PATH = k8s.io/code-generator
CRD_DIR = deploy/crds

.PHONY: go-build
go-build: controller-gen go-mod
	./hack/update-codegen.sh
	./hack/update-codegen-weblogic.sh
	./hack/update-codegen-coherence.sh

	$(CONTROLLER_GEN) object:headerFile=hack/boilerplate.go.txt paths=./pkg/...
	$(CONTROLLER_GEN) crd:crdVersions=v1 output:crd:artifacts:config=deploy/crds paths=./pkg/...
	mv deploy/crds/verrazzano.io_verrazzanomodels.yaml deploy/crds/verrazzano.io_verrazzanomodels_crd.yaml
	mv deploy/crds/verrazzano.io_verrazzanomanagedclusters.yaml deploy/crds/verrazzano.io_verrazzanomanagedclusters_crd.yaml
	mv deploy/crds/verrazzano.io_verrazzanobindings.yaml deploy/crds/verrazzano.io_verrazzanobindings_crd.yaml

	# These crds are generated but not needed
	rm deploy/crds/coherence.oracle.com*
	rm deploy/crds/weblogic.oracle*

	# Add copyright headers to the operator-sdk
	# generated CRDs
	./hack/add-crd-header.sh

.PHONY: go-mod
go-mod:
	$(GO) mod vendor
	$(GO) mod tidy

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

+.PHONY: controller-gen
controller-gen:
	go get sigs.k8s.io/controller-tools/cmd/controller-gen@v0.4.1
ifeq (, $(shell which controller-gen))
CONTROLLER_GEN=$(GOBIN)/controller-gen
else
CONTROLLER_GEN=$(shell which controller-gen)
endif
