package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/sreok/kube-go/docs"
	"github.com/sreok/kube-go/routers"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"log"
)

// @title kube-go
// @version 0.1.0
// @description 基于client-go和gin框架后端api服务
// @termsOfService https://github.com/sreok/kube-go
func main() {
	route := gin.Default()

	route.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	api := route.Group("/api")
	{
		api.GET("/", routers.Index)
	}

	// 指定地址和端口号
	err := route.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
