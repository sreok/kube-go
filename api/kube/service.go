package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/types"
	"github.com/sreok/kube-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
)

// GetServices
// @Summary			获取Service列表
// @Description		获取Service列表
// @Tags			Kubernetes
// @Produce			json
// @Param namespace query string false "命名空间"
// @Router       /api/services [get]
func GetServices(context *gin.Context) {
	namespace := context.DefaultQuery("namespace", "")
	services, err := models.ClientSet().CoreV1().Services(namespace).List(context, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	var svcList []types.SvcSummaryInfo
	for _, svc := range services.Items {
		svcList = append(svcList, types.SvcSummaryInfo{
			Namespace: svc.Namespace,
			Name:      svc.Name,
			Labels:    svc.Labels,
		})
	}
	context.JSON(http.StatusOK, svcList)
}
