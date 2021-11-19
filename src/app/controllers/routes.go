package controllers

import (
	"github.com/bytesfield/golang-gin-auth-service/src/app/middlewares"
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	gin "github.com/gin-gonic/gin"
)

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.GET("/", s.Home)

	api := s.Router.Group("/api/v1")
	{
		// // Auth Route
		api.POST("/login", s.Login)
		api.POST("/token/refresh", middlewares.AuthMiddleware(), s.RefreshToken)
		// //Users routes
		api.POST("/register", s.Register)
		api.GET("/users", middlewares.AuthMiddleware(), s.GetUsers)
		api.GET("/users/:id", s.GetUser)
		api.PUT("/users/:id", middlewares.AuthMiddleware(), s.UpdateUser)
		api.POST("/users/:id", middlewares.AuthMiddleware(), s.DeleteUser)

	}

	//Not Found Route
	s.Router.NoRoute(func(c *gin.Context) {
		responses.NotFound(c, "Not Found")
	})
}
