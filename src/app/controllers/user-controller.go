package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	userRepository "github.com/bytesfield/golang-gin-auth-service/src/app/repositories"
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	gin "github.com/gin-gonic/gin"
)

func (server *Server) GetUsers(ctx *gin.Context) {
	user := models.User{}

	userRepo := userRepository.New(&user)

	users, err := userRepo.FindAllUsers(server.DB)

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

	userRepo := userRepository.New(&user)

	userProfile, err := userRepo.FindUserByID(server.DB, uint32(userId))

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

	err = user.Validate("update")

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err)
		return
	}

	userRepo := userRepository.New(&user)

	err = json.Unmarshal(body, &user)

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	user.Prepare()

	err = user.Validate("update")

	if err != nil {
		responses.ValidationError(ctx, "Validation Error", err.Error())
		return
	}

	updatedUser, err := userRepo.UpdateUser(server.DB, uint32(uid))

	if err != nil {
		responses.ServerError(ctx, "Something went wrong", err.Error())
		return
	}

	responses.Ok(ctx, "User updated successfully", updatedUser)
}

func (server *Server) DeleteUser(ctx *gin.Context) {

	id := ctx.Param("id")

	user := models.User{}

	userRepo := userRepository.New(&user)

	uid, err := strconv.ParseUint(id, 10, 32)

	if err != nil {
		responses.BadRequest(ctx, "Id not integer", err.Error())
		return
	}

	_, err = userRepo.DeleteUser(server.DB, uint32(uid))

	if err != nil {
		responses.ServerError(ctx, "Something went wrong", err.Error())
		return
	}

	responses.Ok(ctx, "User deleted successfully")
}
