package types

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type PodSummaryInfo struct {
	Namespace string
	Name      string
	Status    corev1.PodPhase
	IPs       []corev1.PodIP
	Labels    map[string]string
}

type SvcSummaryInfo struct {
	Namespace string
	Name      string
	Labels    map[string]string
}

type DeploySummaryInfo struct {
	Namespace string
	Name      string
	Labels    map[string]string
}

type NodeSummaryInfo struct {
	Name   string
	Labels map[string]string
}

type NamespaceSummaryInfo struct {
	Name   string
	Labels map[string]string
}

type ApiGroupInfo struct {
	GroupVersion schema.GroupVersion
	ResourceName []metav1.APIResource
}

type VMInfo struct {
	Name      string
	Namespace string
	Status    string
	Running   bool
	Template  interface{}
	Labels    map[string]string
}

type VMIInfo struct {
	Name      string
	Namespace string
	Status    interface{}
	Template  interface{}
	Labels    map[string]string
}
