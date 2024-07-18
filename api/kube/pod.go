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

// GetPods
// @Summary			获取pod列表
// @Description		获取pod列表
// @Tags			Kubernetes
// @Produce			json
// @Param namespace query string false "命名空间"
// @Router       /api/pod [get]
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

// PostPods
// @Summary      获取APIGroups信息
// @Description  获取APIGroups信息
// @Tags         Kubernetes
// @Router       /api/pod [post]
func PostPods(context *gin.Context) {
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
	context.JSON(http.StatusOK, apiInfo)
}

// DeletePods
// @Summary      获取APIGroups信息
// @Description  获取APIGroups信息
// @Tags         Kubernetes
// @Router       /api/pod [delete]
func DeletePods(context *gin.Context) {
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
	context.JSON(http.StatusOK, apiInfo)
}

// PutPods
// @Summary      获取APIGroups信息
// @Description  获取APIGroups信息
// @Tags         Kubernetes
// @Router       /api/pod [put]
func PutPods(context *gin.Context) {
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
	context.JSON(http.StatusOK, apiInfo)
}
