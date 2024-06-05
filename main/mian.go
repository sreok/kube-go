package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sreok/kube-go/routers/api"
	"log"
)

func main() {
	router := gin.Default()

	router.GET("/", api.Index)
	// 指定地址和端口号
	err := router.Run(":8080")
	if err != nil {
		log.Println(err)
		return
	}
}
