# Verrazzano CRD generator

To generate Kubernetes custom resource definitions (CRD) for the `verrazzano-operator` [project](https://github.com/verrazzano/verrazzano-operator),
use the Verrazzano CRD generator.
## Prerequisites

[Operator SDK](https://github.com/operator-framework/operator-sdk) version v0.18.1 must be installed on your system.

## How to build

To generate the CRD, run the following command:

```
make go-build
```

## How to update the copyright header in generated files

The build process generates several files which include a copyright header. To update the copyright header of the
generated files, modify the contents of the following files and build the project:
* `hack/boilerplate.go.txt` - The contents of this file are added to files generated in `pkg/apis`.
* `hack/boilerplate.yaml.txt` - The contents of this file are added to files generated in `deploy/crds`.
* `hack/custom-header.txt` - The contents of this file are added to files generated in `pkg/client`, `pkg/clientcoherence`, and `pkg/clientwks`.
