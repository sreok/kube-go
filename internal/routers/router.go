package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/internal/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

// GetNamespaces
// @Summary      获取命名空间
// @Description  获取命名空间
// @Tags         Kubernetes
// @Router       /api/namespaces [get]
func GetNamespaces(context *gin.Context) {
	namespaces, err := models.KubeClient().CoreV1().Namespaces().List(context, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	for _, namespace := range namespaces.Items {
		log.Println(namespace.Name)
	}
}

// GetPods
// @Summary      获取pod列表
// @Description  获取pod列表
// @Tags         Kubernetes
// @Router       /api/pods [get]
func GetPods(context *gin.Context) {
	pods, err := models.KubeClient().CoreV1().Pods("default").List(context, metav1.ListOptions{})
	if err != nil {
		log.Println(err.Error())
	}
	for _, pod := range pods.Items {
		log.Println(pod.Name)
	}
}
