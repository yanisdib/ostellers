package artbook

import "github.com/gin-gonic/gin"

func NewArtbookRoutes(router *gin.Engine) {

	router.POST("/artbook", Create())
	router.GET("/artbooks/:id", GetByID())
	router.DELETE("/artbooks/:id", DeleteByID())
	router.PUT("/artbooks/:id", UpdateByID())

}
