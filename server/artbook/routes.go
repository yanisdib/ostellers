package artbook

import "github.com/gin-gonic/gin"

func NewArtbookRoutes(router *gin.Engine) {
	router.POST("/artbook", Create())
	router.GET("/artbook/:id", GetByID())
}
