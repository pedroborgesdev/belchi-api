package main

import (
	"bellchi/config"
	"bellchi/database"
	"bellchi/debug"
	"bellchi/logger"
	"bellchi/middlewares"
	"bellchi/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	debug.LoadDebugConfig()

	logger.Log("INFO", "Application has been started", []logger.LogDetail{})

	config.LoadAppConfig()

	database.InitDB()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(
		middlewares.CORSMiddleware(),
	)

	routes.SetupRoutes(router)

	router.Run(":" + config.AppConfig.HTTPPort)
}
