package cmd

import (
	"github.com/SaroarShahan/event-management/api/handlers"
	"github.com/SaroarShahan/event-management/infra/database"
)

func migration() {
	err := database.DB.AutoMigrate(&handlers.User{}, &handlers.Event{}, &handlers.Registration{})

	if err != nil {
		panic("Database migration failed: " + err.Error())
	}
}