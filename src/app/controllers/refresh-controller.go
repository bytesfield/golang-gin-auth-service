package controllers

import (
	"fmt"
	"strconv"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	"github.com/bytesfield/golang-gin-auth-service/src/app/services"
	gin "github.com/gin-gonic/gin"
)

func (server *Server) RefreshToken(ctx *gin.Context) {

	id, err := services.GetTokenID(ctx)
	fmt.Println(err)
	if err != nil {
		responses.Unauthorized(ctx, "Unauthorized")
		return
	}

	userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", id), 10, 32)

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

	token, err := services.RefreshToken(ctx)

	if err != nil {
		responses.BadRequest(ctx, err.Error())
		return
	}

	userData := map[string]interface{}{"token": token, "user": userProfile}

	responses.Ok(ctx, "Token refreshed successfully", userData)
}
