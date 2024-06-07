package routers

import (
	_ "context"
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/internal/models"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

// GetNamespacesList
// @Summary      Show an account
// @Description  获取命名空间
// @Tags         获取命名空间
// @Router       /api/namespaces [get]
func GetNamespacesList(context *gin.Context) {
	clientSet, err := models.KubeClient()
	if err != nil {
		log.Panic(err.Error())
	}
	var namespaces *v1.NamespaceList
	namespaces, err = clientSet.CoreV1().Namespaces().List(context, metav1.ListOptions{})
	if err != nil {
		log.Panic(err)
	}
	for _, namespace := range namespaces.Items {
		log.Println(namespace.Name)
		log.Println(namespace.Labels)
	}
}
