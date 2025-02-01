package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"my-rest-api/db"
	_ "my-rest-api/docs" // This is important to import generated docs
	"my-rest-api/routes"
)

// @title My REST API
// @version 1.0
// @description This is a sample server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":8080")
}
