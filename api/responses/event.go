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

	userId := context.GetInt64("userId")
	event.UserID = &userId

	if err := event.SaveEventsHandler(); err != nil {
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
		"data": event,
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

	userId := context.GetInt64("userId")
	event, err := handlers.GetEventHandler(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Event not found",
			"data": nil,
		})
		return
	}

	if event.UserID == nil || *event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{
			"status": false,
			"message": "You are not authorized to update this event",
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

	updatedEvent.ID = uint(id)

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

	userId := context.GetInt64("userId")
	event, err := handlers.GetEventHandler(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Event not found",
			"data": nil,
		})
		return
	}

	if event.UserID == nil || *event.UserID != userId {
		context.JSON(http.StatusForbidden, gin.H{
			"status": false,
			"message": "You are not authorized to delete this event",
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

func RegisterEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	event, err := handlers.GetEventHandler(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"message": "Event not found",
			"data": nil,
		})
		return
	}

	if err := event.RegisterEventHandler(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to register for event",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Successfully registered for the event!",
		"data": nil,
	})
}

func DeleteEventRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"message": "Invalid event ID",
			"data": nil,
		})
		return
	}

	var event handlers.Event
	event.ID = uint(eventId)

	if err := event.DeleteEventRegistrationHandler(userId); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": "Failed to delete event registration",
			"data": nil,
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Successfully deleted event registration!",
		"data": nil,
	})
}