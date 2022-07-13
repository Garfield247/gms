package router

import (
	"github.com/Garfield247/gms.git/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(config.ServerSetting.RunMode)

	img := r.Group("/image")
	{
		img.GET("/")
	}
}
