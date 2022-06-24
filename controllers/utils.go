package controllers

import (
	corev1 "k8s.io/api/core/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/util/conditions"
)

func IsTrue(from conditions.Getter, t clusterv1.ConditionType) bool {
	if c := Get(from, t); c != nil {
		return c.Status == corev1.ConditionTrue
	}
	return false
}

func Get(from conditions.Getter, t clusterv1.ConditionType) *clusterv1.Condition {
	for _, condition := range from.GetConditions() {
		if condition.Type == t {
			return &condition
		}
	}
	return nil
}
