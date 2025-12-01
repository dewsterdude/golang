package routes

import (
	"log"
	"net/http"
	"strconv"

	"example.com/REST-API/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) { // structure to frame up response
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve events"}) // 500
		return
	}
	context.JSON(http.StatusOK, events)

	//	context.JSON(http.StatusOK, gin.H{"message": "Hello!"}) // JSON is default for REST APIs, gin.H is a map function
	// "message is the key, Hello! is the value"
}

func getEvent(context *gin.Context) { // convert base=10, 64-bit integer
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the dynamic parameter from the URL -- as a STRING
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	// Old code moved to Auth.go
	/*	token := context.Request.Header.Get("Authorization") // commonly used to transmit tokens
		if token == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
			return
		}

		userID, err := utils.VerifyToken(token) // during this verification, the claims of data including user id is used

		if err != nil { // theres an error
			context.JSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
			return
		}
	*/
	var event models.Event

	err := context.ShouldBindJSON(&event) // bind the incoming JSON to the event struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error binding json in createEvent"})
		return
	}
	userId := context.GetInt64("userId") // fetch value from gin context
	event.UserID = userId

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event}) // Code 201 - New Resource

	// gin will take a look at the incomming JSON body and store it in your structure
	// clients must send requests that matches the structure

}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the dynamic parameter from the URL -- as a STRING
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}
	//	fmt.Println("update event ID: ", eventId)

	userID := context.GetInt64("userId") // fetch value from gin context
	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event ID"})
		return
	}
	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return

	}

	var updatedEvent models.Event
	err = context.ShouldBindJSON(&updatedEvent) // bind the incoming JSON to the event struct

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request when updating"})
		return
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event ID"})
		log.Println("error on update event id in updateEvent()", err)
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // get the dynamic parameter from the URL -- as a STRING
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID for delete"})
		return
	}

	userID := context.GetInt64("userId") // fetch value from gin context

	// check if event actually exists
	deleteEvent, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event ID for dlete"})
		return
	}

	if deleteEvent.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
		return
	}

	// execute delete
	err = deleteEvent.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event ID"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
