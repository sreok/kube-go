package kube

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/mysql"
	"github.com/sreok/kube-go/api/types"
	"github.com/sreok/kube-go/models"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"net/http"
)

// GetAPIGroups
// @Summary      获取APIGroups信息
// @Description  获取APIGroups信息
// @Tags         Kubernetes
// @Router       /api/groups [get]
func GetAPIGroups(context *gin.Context) {
	// 获取所有分组和资源数据
	_, APIResourceListSlice, err := models.DiscoveryClient().ServerGroupsAndResources()
	if err != nil {
		panic(err.Error())
	}
	var apiInfo []types.ApiGroupInfo
	for _, singleAPIResourceList := range APIResourceListSlice {
		// GroupVersion是个字符串，例如"apps/v1"
		groupVersionStr := singleAPIResourceList.GroupVersion
		// ParseGroupVersion方法将字符串转成数据结构
		gv, err := schema.ParseGroupVersion(groupVersionStr)
		if err != nil {
			panic(err.Error())
		}
		apiInfo = append(apiInfo, types.ApiGroupInfo{
			GroupVersion: gv,
			ResourceName: singleAPIResourceList.APIResources,
		})
	}
	mysql.InsertApiGroups()
	context.JSON(http.StatusOK, apiInfo)
}
