package controllers

import (
	"fmt"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	userRepository "github.com/bytesfield/golang-gin-auth-service/src/app/repositories"
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

	user := models.User{}

	userRepo := userRepository.New(&user)

	userProfile, err := userRepo.FindUserByID(server.DB, uint32(id))

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
