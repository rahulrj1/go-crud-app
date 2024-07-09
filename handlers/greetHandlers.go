package handlers

import (
	"go-crud-app/model"

	"github.com/gin-gonic/gin"
)

func Greet(cxt *gin.Context) {
	cxt.JSON(200, gin.H{"message": "Hello, how are you?"})
}

func GreetToName(cxt *gin.Context) {
	var user model.User
	err := cxt.ShouldBindJSON(&user)
	if err != nil {
		cxt.JSON(400, gin.H{"message": "Error binding JSON data"})
	}

	cxt.JSON(200, gin.H{"message": "Hello " + user.Name + ", how are you?"})
}
