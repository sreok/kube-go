package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/types"
	"github.com/sreok/kube-go/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"net/http"
)

// GetNamespaces
// @Summary      获取命名空间
// @Description  获取命名空间
// @Tags         Kubernetes
// @Router       /api/namespaces [get]
func GetNamespaces(context *gin.Context) {
	namespaces, err := models.ClientSet().CoreV1().Namespaces().List(context, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	var nsList []types.NamespaceSummaryInfo
	for _, pod := range namespaces.Items {
		nsList = append(nsList, types.NamespaceSummaryInfo{
			Name:   pod.Name,
			Labels: pod.Labels,
		})
	}
	context.JSON(http.StatusOK, nsList)
}
