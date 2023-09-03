package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stivenviedman/go-mongo-starter/src/config"
	"github.com/stivenviedman/go-mongo-starter/src/dummy"

	docs "github.com/stivenviedman/go-mongo-starter/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Init() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/v1"

	v1 := r.Group("/v1")

	dummy.RegisterRoutes(v1)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL(config.Env.APP_BASE_URL+"/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	r.Run(":3000")
}
