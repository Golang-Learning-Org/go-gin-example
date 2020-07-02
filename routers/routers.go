package routers

import (
	"github.com/evanxzj/go-gin-example/middleware/jwt"
	"github.com/evanxzj/go-gin-example/pkg/setting"
	"github.com/evanxzj/go-gin-example/routers/api"
	v1 "github.com/evanxzj/go-gin-example/routers/api/v1"

	"github.com/gin-gonic/gin"

	_ "github.com/evanxzj/go-gin-example/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Auth
	r.GET("/auth", api.GetAuth)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// Tags API
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		// Articles API
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
