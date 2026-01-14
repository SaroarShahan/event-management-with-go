package cmd

import (
	"strconv"

	"github.com/SaroarShahan/event-management/config"
	"github.com/SaroarShahan/event-management/infra/database"
	"github.com/SaroarShahan/event-management/routes"
	"github.com/gin-gonic/gin"
)

func Server() {
	database.InitDB()
	config := config.GetConfig()
	app := gin.Default()

	routes.RegisterEventRoutes(app)
	
	app.Run(":" + strconv.Itoa(config.HttpPort))
}