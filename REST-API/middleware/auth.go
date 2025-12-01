package middleware

import (
	"net/http"

	"example.com/REST-API/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {

	token := context.Request.Header.Get("Authorization") // commonly used to transmit tokens

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	userID, err := utils.VerifyToken(token) // during this verification, the claims of data including user id is used

	if err != nil { // theres an error
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized."})
		return
	}

	context.Set("userId", userID) // attach new data to the context

	context.Next() // this ensures the next request handler in line executes correctly
}
