package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/types"
	"github.com/sreok/kube-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"log"
	"net/http"
)

// GetVMs
// @Summary			获取vm列表
// @Description		获取vm列表
// @Tags			Kubernetes
// @Produce			json
// @Param namespace query string false "命名空间"
// @Router       	/api/vm/list [get]
func GetVMs(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "")
	groupSource := schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachines"}
	vms, err := models.DynamicClient().Resource(groupSource).Namespace(namespace).List(context, metav1.ListOptions{})
	if err != nil || vms == nil {
		log.Println(err.Error())
	}
	var vmList []types.VMInfo
	for _, vm := range vms.Items {
		vmList = append(vmList, types.VMInfo{
			Name:      vm.GetName(),
			Namespace: vm.GetNamespace(),
			Labels:    vm.GetLabels(),
			Status:    vm.UnstructuredContent()["status"].(map[string]interface{})["printableStatus"].(string),
			Running:   vm.UnstructuredContent()["spec"].(map[string]interface{})["running"].(bool),
			Template:  vm.UnstructuredContent()["spec"].(map[string]interface{})["template"],
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"vmList": vmList,
	})
}

// GetVMIs
// @Summary			获取vmi列表
// @Description		获取vmi列表
// @Tags			Kubernetes
// @Produce			json
// @Param namespace query string false "命名空间"
// @Router       	/api/vmi/list [get]
func GetVMIs(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "")
	groupSource := schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachineinstances"}
	vms, err := models.DynamicClient().Resource(groupSource).Namespace(namespace).List(context, metav1.ListOptions{})
	if err != nil || vms == nil {
		log.Println(err.Error())
	}
	var vmiList []types.VMIInfo
	for _, vmi := range vms.Items {
		vmiList = append(vmiList, types.VMIInfo{
			Name:      vmi.GetName(),
			Namespace: vmi.GetNamespace(),
			Labels:    vmi.GetLabels(),
			Status:    vmi.UnstructuredContent()["status"],
			Template:  vmi.UnstructuredContent()["spec"],
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"vmiList": vmiList,
	})
}
