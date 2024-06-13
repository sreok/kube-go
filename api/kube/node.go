package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/types"
	"github.com/sreok/kube-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
)

// GetNodes
// @Summary      获取node列表
// @Description  获取node列表
// @Tags         Kubernetes
// @Router       /api/nodes [get]
func GetNodes(context *gin.Context) {
	nodes, err := models.ClientSet().CoreV1().Nodes().List(context, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	var nodeList []types.NodeSummaryInfo
	for _, svc := range nodes.Items {
		nodeList = append(nodeList, types.NodeSummaryInfo{
			Name:   svc.Name,
			Labels: svc.Labels,
		})
	}
	context.JSON(http.StatusOK, nodeList)
}
