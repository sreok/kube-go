package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/types"
	"github.com/sreok/kube-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
)

// GetDeployments
// @Summary			获取Deployments列表
// @Description		获取Deployments列表
// @Tags			Kubernetes
// @Produce			json
// @Param namespace query string false "命名空间"
// @Router       /api/deployments [get]
func GetDeployments(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "")
	deployments, err := models.ClientSet().AppsV1().Deployments(namespace).List(context, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	var deployList []types.DeploySummaryInfo
	for _, svc := range deployments.Items {
		deployList = append(deployList, types.DeploySummaryInfo{
			Namespace: svc.Namespace,
			Name:      svc.Name,
			Labels:    svc.Labels,
		})
	}
	context.JSON(http.StatusOK, deployList)
}
