package main

import (
	"github.com/gin-gonic/gin"

	"github.com/SaroarShahan/event-management/infra/database"
	"github.com/SaroarShahan/event-management/routes"
)

func main() {
	database.InitDB()
	app := gin.Default()

	routes.RegisterEventRoutes(app)

	app.Run(":8080")
}
