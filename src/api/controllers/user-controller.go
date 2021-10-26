package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bytesfield/golang-gin-auth-service/src/api/models"
	"github.com/bytesfield/golang-gin-auth-service/src/api/responses"
	"github.com/bytesfield/golang-gin-auth-service/src/api/services"
	gin "github.com/gin-gonic/gin"
)

func (server *Server) CreateUser(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		responses.ValidationError(c, "Validation Error", err)
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ValidationError(c, "Validation Error", err)
		return
	}

	user.Prepare()

	err = user.Validate("")

	if err != nil {
		responses.ValidationError(c, "Validation Error", err)
		return
	}

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		//formattedError := FormatError .FormatError(err.Error())
		responses.ServerError(c, "Validation Error", err.Error())

		return
	}

	c.Header("Location", fmt.Sprintf("%s%s/%d", c.Request.Host, c.Request.RequestURI, userCreated.ID))

	responses.CreatedResponse(c, "User added successfully", userCreated)

}

func (server *Server) GetUsers(c *gin.Context) {
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	fmt.Println(users)

	responses.Ok(c, "User added successfully", users)
}

func (server *Server) GetUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		responses.BadRequest(c, "Error occurred", err)
		return
	}

	user := models.User{}

	userProfile, err := user.FindUserByID(server.DB, uint32(userId))

	if err != nil {
		responses.NotFound(c, "User not found", err)
		return
	}

	responses.Ok(c, "User retrieved successfully", userProfile)
}

func (server *Server) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		responses.BadRequest(c, "Error occurred", err)
		return
	}

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		responses.ValidationError(c, "Validation Error", err)
		return
	}
	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ValidationError(c, "Validation Error", err)
		return
	}

	tokenID, err := services.ExtractTokenID(c.Request)

	if err != nil {
		responses.Unauthorized(c, "Unauthorized")
		return
	}

	if tokenID != uint32(uid) {
		responses.Unauthorized(c, "Unauthorized")
		return
	}

	user.Prepare()

	err = user.Validate("update")

	if err != nil {
		responses.ValidationError(c, "Validation Error", err)
		return
	}

	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))

	if err != nil {
		// formattedError := formaterror.FormatError(err.Error())

		responses.ServerError(c, "Something went wrong", err)
		return
	}

	responses.Ok(c, "User updated successfully", updatedUser)
}

func (server *Server) DeleteUser(c *gin.Context) {

	id := c.Param("id")

	user := models.User{}

	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		responses.BadRequest(c, "Id not integer", err)
		return
	}

	tokenID, err := services.ExtractTokenID(c.Request)

	if err != nil {
		responses.Unauthorized(c, "Unauthorized", err)
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.Unauthorized(c, "Unauthorized", err)
		return
	}

	_, err = user.DeleteAUser(server.DB, uint32(uid))

	if err != nil {
		responses.ServerError(c, "Something went wrong", err)
		return
	}

	responses.NoContent(c, "No Content", err)
}
