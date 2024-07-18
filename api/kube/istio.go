package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
)

// GetIstioGateways
// @Summary      获取gateway信息
// @Description  获取gateway信息
// @Tags         Kubernetes
// @Param namespace query string false "命名空间"
// @Router       /api/gateways [get]
func GetIstioGateways(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "")
	// dynamicClient的唯一关联方法所需的入参
	groupSource := schema.GroupVersionResource{Group: "networking.istio.io", Version: "v1beta1", Resource: "gateways"}
	gws, err := models.DynamicClient().Resource(groupSource).Namespace(namespace).List(context, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	for _, gw := range gws.Items {
		log.Println(gw.GetName())
	}
	//podList := &corev1.PodList{}
	//err = runtime.DefaultUnstructuredConverter.FromUnstructured(obj.UnstructuredContent(), podList)
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Printf("namespace\t status\t\t name\n")
	//for _, d := range podList.Items {
	//	fmt.Printf("%v\t %v\t %v\n",
	//		d.Namespace,
	//		d.Status.Phase,
	//		d.Name)
	//}
}
