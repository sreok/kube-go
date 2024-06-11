package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sreok/kube-go/docs/swagger"
	"github.com/sreok/kube-go/internal/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title kube-go
// @version 0.1.0
// @description 基于client-go和gin框架后端api服务
// @termsOfService https://github.com/sreok/kube-go
func main() {
	r := gin.Default()

	// swagger文档
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		api.GET("/namespaces", routers.GetNamespaces)
		api.GET("/pods", routers.GetPods)
	}

	// 指定地址和端口号
	err := r.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
