package kube

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/models"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GetGateways
// @Summary      获取gateway信息
// @Description  获取gateway信息
// @Tags         Kubernetes
// @Param namespace query string false "命名空间"
// @Router       /api/gateways [get]
func GetGateways(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "")
	// dynamicClient的唯一关联方法所需的入参
	groupSource := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	obj, err := models.DynamicClient().Resource(groupSource).Namespace(namespace).List(context, metav1.ListOptions{Limit: 100})
	if err != nil {
		panic(err.Error())
	}
	podList := &corev1.PodList{}
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), podList)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("namespace\t status\t\t name\n")
	for _, d := range podList.Items {
		fmt.Printf("%v\t %v\t %v\n",
			d.Namespace,
			d.Status.Phase,
			d.Name)
	}
}
