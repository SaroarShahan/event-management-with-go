package responses

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SaroarShahan/event-management/api/handlers"
	"github.com/gin-gonic/gin"
)


func CreateEvent(context *gin.Context) {
	var event handlers.Event

	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid request payload",
			"data": nil,
		})
		return
	}

	createdEvent, err := handlers.CreateEventHandler(event)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to create event",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"status": true,
		"message": "Event has been created successfully!",
		"data": createdEvent,
	})
}

func GetEvents(context *gin.Context) {
	events, err := handlers.GetAllEventsHandler()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to fetch events",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Events has been fetched successfully!",
		"data": events,
	})
}

func GetEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	event, err := handlers.GetEventHandler(id)

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

func UpdateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	if _, err = handlers.GetEventHandler(id); err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Event not found",
			"data": nil,
		})
		return
	}

	var updatedEvent handlers.Event

	if err := context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid request payload",
			"data": nil,
		})
		return
	}

	updatedEvent.ID = id

	err = handlers.UpdateEventHandler(updatedEvent)
	
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

func DeleteEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	_, err = handlers.GetEventHandler(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Event not found",
			"data": nil,
		})
		return
	}

	err = handlers.DeleteEventHandler(id)

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