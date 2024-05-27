package artbook

import "github.com/gin-gonic/gin"

func NewArtbookRoutes(router *gin.Engine) {
	router.POST("v1/artbook", Create())
	router.GET("v1/artbooks", GetAll())
	router.GET("v1/artbooks/:id", GetByID())
	router.DELETE("v1/artbooks/:id", DeleteByID())
	router.PUT("v1/artbooks/:id", UpdateByID())
}
