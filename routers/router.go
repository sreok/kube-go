package routers

import "github.com/gin-gonic/gin"

// Index godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         获取命名空间
// @Accept       json
// @Produce      json
// @Router       /namespace [get]
func Index(context *gin.Context) {
	context.JSON(200, gin.H{
		"code":    200,
		"success": true,
	})
}
