package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Build(db *gorm.DB) *gin.Engine {
	route := gin.Default()
	route.Use(gin.Recovery())
	route.SetTrustedProxies([]string{"127.0.0.1", "localhost"})
	route.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"app_name": "deltomed_ticketing",
			"version":  "1.0.0",
		})
	})
	route.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "route not found",
			"data":    nil,
		})
	})
	return route
}
