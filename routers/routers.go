package routers

import (
	"github.com/evanxzj/go-gin-example/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	return r
}
