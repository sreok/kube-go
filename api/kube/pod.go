package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/types"
	"github.com/sreok/kube-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
)

// GetPods
// @Summary			获取pod列表
// @Description		获取pod列表
// @Tags			Kubernetes
// @Produce			json
// @Param namespace query string false "命名空间"
// @Router       /api/pod/list [get]
func GetPods(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "")
	pods, err := models.ClientSet().CoreV1().Pods(namespace).List(context, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	var podList []types.PodSummaryInfo
	for _, pod := range pods.Items {
		podList = append(podList, types.PodSummaryInfo{
			Namespace: pod.Namespace,
			Name:      pod.Name,
			Status:    pod.Status.Phase,
			IPs:       pod.Status.PodIPs,
			Labels:    pod.Labels,
		})
	}
	context.JSON(http.StatusOK, podList)
}
