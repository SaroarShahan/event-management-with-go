package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SaroarShahan/event-management/models"
	"github.com/gin-gonic/gin"
)


func createEvent(context *gin.Context) {
	var event models.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid request payload",
			"data": nil,
		})
		fmt.Println("Error binding JSON:", err)
		return
	}

	createdEvent, err := models.CreateEvent(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to create event",
			"data": nil,
		})
		fmt.Println("Error creating event:", err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status": true,
		"message": "Event has been created successfully!",
		"data": createdEvent,
	})
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to fetch events",
			"data": nil,
		})
		fmt.Println("Error fetching events:", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Events has been fetched successfully!",
		"data": events,
	})
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	event, err := models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to fetch event",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Event has been fetched successfully!",
		"data": event,
	})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	if _, err = models.GetEvent(id); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Event not found",
			"data": nil,
		})
		return
	}

	var updatedEvent models.Event

	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid request payload",
			"data": nil,
		})
		return
	}

	updatedEvent.ID = id

	err = models.UpdateEvent(updatedEvent)
	
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to update event",
			"data": nil,
		})
		fmt.Println("Error updating event:", err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Event has been updated successfully!",
		"data": updatedEvent,
	})
}

func deleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	_, err = models.GetEvent(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Event not found",
			"data": nil,
		})
		return
	}

	err = models.DeleteEvent(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to delete event",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Event has been deleted successfully!",
		"data": nil,
	})
}