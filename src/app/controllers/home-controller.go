package controllers

import (
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	gin "github.com/gin-gonic/gin"
)

func (server *Server) Home(c *gin.Context) {

	responses.Ok(c, "Welcome To Golang Gin Auth Service API")
}
