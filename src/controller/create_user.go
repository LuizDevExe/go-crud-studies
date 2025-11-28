package controller

import (
	"fmt"
	"log"

	"github.com/LuizDevExe/go-crud-studies/src/configuration/validation"
	"github.com/LuizDevExe/go-crud-studies/src/controller/model/request"
	"github.com/LuizDevExe/go-crud-studies/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	var userRequest request.CreateUserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error:%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println(userRequest)

	userResponse := response.UserResponse{
		ID:    "teste",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	c.JSON(201, userResponse)
}
