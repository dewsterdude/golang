package routes

import (
	"net/http"

	"example.com/REST-API/models"
	"example.com/REST-API/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user) // bind the incoming JSON to the event struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error binding json in signup"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Usercreated!", "user": user}) // Code 201 - New Resource
	// gin will take a look at the incomming JSON body and store it in your structure
	// clients must send requests that matches the structure
}

func login(context *gin.Context) {
	var user models.User
	//	fmt.Println("Start Login")

	err := context.ShouldBindJSON(&user) // bind the incoming JSON to the event struct populate struct
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "error binding json to login struct"})
		return
	}
	//	fmt.Println("Before Validate Credentials")

	err = user.ValidateCredentials()
	//	fmt.Println("After Validate Credentials")

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not Authenticate User."})
	}
	//	fmt.Println("before generating token")
	token, err := utils.GenerateToken(user.Email, user.ID)
	//	fmt.Println("after generating token")

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Authenticate User2"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful!", "token": token})

}
