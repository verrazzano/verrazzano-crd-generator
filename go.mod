module github.com/verrazzano/verrazzano-crd-generator

require (
	github.com/go-openapi/spec v0.19.0
	github.com/operator-framework/operator-sdk v0.12.0
	github.com/spf13/pflag v1.0.3
	github.com/stretchr/testify v1.4.0
	gopkg.in/yaml.v2 v2.2.4
	istio.io/api v0.0.0-20200107183329-ed4b507c54e1
	k8s.io/api v0.15.7
	k8s.io/apimachinery v0.15.7
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.15.7
	k8s.io/gengo v0.0.0-20191010091904-7fa3014cb28f
	k8s.io/kube-openapi v0.0.0-20190918143330-0270cf2f1c1d
	sigs.k8s.io/controller-runtime v0.3.0
)

// Pinned to kubernetes-0.15.7
replace (
	gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.2.2
	k8s.io/api => k8s.io/api v0.15.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.15.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.15.7
	k8s.io/apiserver => k8s.io/apiserver v0.15.7
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.15.7
	k8s.io/client-go => k8s.io/client-go v0.15.7
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.15.7
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.15.7
	k8s.io/code-generator => k8s.io/code-generator v0.15.7
	k8s.io/component-base => k8s.io/component-base v0.15.7
	k8s.io/cri-api => k8s.io/cri-api v0.15.7
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.15.7
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.15.7
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.15.7
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.15.7
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.15.7
	k8s.io/kubelet => k8s.io/kubelet v0.15.7
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.15.7
	k8s.io/metrics => k8s.io/metrics v0.15.7
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.15.7
)

// Pinned to istio 1.4.6
replace istio.io/api => istio.io/api v0.0.0-20200107183329-ed4b507c54e1

go 1.13