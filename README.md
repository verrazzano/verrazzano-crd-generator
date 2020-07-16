# Verrazzano CRD generator

The Verrazzano CRD generator generates Kubernetes custom resource definitions (CRD)
used by the Verrazzano `verrazzano-operator` [project](https://github.com/verrazzano/verrazzano-operator).

## Prerequisites

The following software must be installed on your system.
* [operator-sdk](https://github.com/operator-framework/operator-sdk) version v0.18.1

## How to build

Run the following command to generate the CRD:

```
make go-build
```

## How to update the copyright header in generated files

The build process generates several files which include a copyright header. To update the copyright header of
generated files, modify the contents of these following files and build the project:
* `hack/boilerplate.go.txt` - The contents of this file are added to files generated in `pkg/apis`.
* `hack/boilerplate.yaml.txt` - The contents of this file are added to files generated in `deploy/crds`.
* `hack/custom-header.txt` - The contents of this file are added to files generated in `pkg/client`, `pkg/clientcoherence`, `pkg/clientistio`, and `pkg/clientwks`.
