package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	"github.com/bytesfield/golang-gin-auth-service/src/app/services"
	gin "github.com/gin-gonic/gin"
)

func (server *Server) Register(ctx *gin.Context) {

	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	user.Prepare()

	err = user.Validate("")

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	userCreated, err := user.SaveUser(server.DB)

	if err != nil {
		//formattedError := FormatError .FormatError(err.Error())
		responses.ServerError(ctx, "Validation Error", err.Error())

		return
	}

	ctx.Header("Location", fmt.Sprintf("%s%s/%d", ctx.Request.Host, ctx.Request.RequestURI, userCreated.ID))

	responses.CreatedResponse(ctx, "Registration successfully", userCreated)

}

func (server *Server) GetUsers(ctx *gin.Context) {
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)

	if err != nil {
		responses.ServerError(ctx, "Something went wrong", err.Error())
		return
	}
	fmt.Println(users)

	responses.Ok(ctx, "Users retrieved successfully", users)
}

func (server *Server) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		responses.BadRequest(ctx, "Error occurred", err.Error())
		return
	}

	user := models.User{}

	userProfile, err := user.FindUserByID(server.DB, uint32(userId))

	if err != nil {
		responses.NotFound(ctx, "User not found", err.Error())
		return
	}

	responses.Ok(ctx, "User retrieved successfully", userProfile)
}

func (server *Server) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")

	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		responses.BadRequest(ctx, "Error occurred", err.Error())
		return
	}

	body, err := ioutil.ReadAll(ctx.Request.Body)

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}
	user := models.User{}

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	tokenID, err := services.GetTokenID(ctx)

	if err != nil {
		responses.Unauthorized(ctx, "Unauthorized")
		return
	}

	if tokenID != uint32(uid) {
		responses.Unauthorized(ctx, "Unauthorized")
		return
	}

	user.Prepare()

	err = user.Validate("update")

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))

	if err != nil {
		// formattedError := formaterror.FormatError(err.Error())

		responses.ServerError(ctx, "Something went wrong", err.Error())
		return
	}

	responses.Ok(ctx, "User updated successfully", updatedUser)
}

func (server *Server) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")

	user := models.User{}

	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		responses.BadRequest(ctx, "Id not integer", err.Error())
		return
	}

	tokenID, err := services.GetTokenID(ctx)

	if err != nil {
		responses.Unauthorized(ctx, "Unauthorized", err.Error())
		return
	}
	if tokenID != 0 && tokenID != uint32(uid) {
		responses.Unauthorized(ctx, "Unauthorized", err.Error())
		return
	}

	_, err = user.DeleteAUser(server.DB, uint32(uid))

	if err != nil {
		responses.ServerError(ctx, "Something went wrong", err.Error())
		return
	}

	responses.NoContent(ctx, "No Content", err.Error())
}
