package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/api/kube"
	_ "github.com/sreok/kube-go/docs/swagger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title kube-go
// @version 0.1.0
// @description 基于client-go和gin框架后端api服务
// @termsOfService https://github.com/sreok/kube-go
func main() {
	route := gin.Default()
	// swagger文档
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routeApi := route.Group("/api")
	{
		routeApi.GET("/namespaces", kube.GetNamespaces)
		pod := routeApi.Group("/pod")
		{
			pod.GET("/list", kube.GetPods)
		}
		routeApi.GET("/nodes", kube.GetNodes)
		routeApi.GET("/groups", kube.GetAPIGroups)
		routeApi.GET("/gateways", kube.GetGateways)
		routeApi.GET("/services", kube.GetServices)
		routeApi.GET("/deployments", kube.GetDeployments)
		vm := routeApi.Group("/vm")
		{
			vm.GET("/list", kube.GetVMs)
		}
		vmi := routeApi.Group("/vmi")
		{
			vmi.GET("/list", kube.GetVMIs)
		}
	}

	err := route.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
