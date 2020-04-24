package subscription

import (
	operatorsv1beta1 "github.com/open-cluster-management/multicloudhub-operator/pkg/apis/operators/v1beta1"
	"github.com/open-cluster-management/multicloudhub-operator/pkg/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// RCM overrides the rcm chart
func RCM(m *operatorsv1beta1.MultiClusterHub, cache utils.CacheSpec) *unstructured.Unstructured {
	sub := &Subscription{
		Name:      "rcm",
		Namespace: m.Namespace,
		Overrides: map[string]interface{}{
			"hubconfig": map[string]interface{}{
				"replicaCount": m.Spec.ReplicaCount,
				"nodeSelector": m.Spec.NodeSelector,
			},
			"global": map[string]interface{}{
				"pullPolicy":      utils.GetImagePullPolicy(m),
				"imagePullSecret": m.Spec.ImagePullSecret,
				"imageRepository": m.Spec.Overrides.ImageRepository,
				"imageTagPostfix": imageSuffix(m),
				"imageOverrides":  cache.ImageOverrides,
			},
		},
	}

	return newSubscription(m, sub)
}