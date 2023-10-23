package app

import (
	"github.com/DrNikita/AstrologerService_TestTask/docs"
	_ "github.com/DrNikita/AstrologerService_TestTask/docs"
	config2 "github.com/DrNikita/AstrologerService_TestTask/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	router := gin.Default()
	router.Use(cors.Default())
	setupRoutes(router)
	router.GET("/swagger/*any", func(context *gin.Context) {
		docs.SwaggerInfo.Host = context.Request.Host
		ginSwagger.CustomWrapHandler(&ginSwagger.Config{URL: "/swagger/doc.json"}, swaggerFiles.Handler)(context)
	})

	config := config2.GetConfigurationInstance()

	err := router.Run(config.AppPort)
	if err != nil {
		log.Fatalf("couldn't run app on port: " + config.AppPort)
	}
}

func setupRoutes(router *gin.Engine) {
}
