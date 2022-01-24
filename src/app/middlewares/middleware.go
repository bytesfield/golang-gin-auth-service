package middlewares

import (
	"net/http"

	"github.com/bytesfield/golang-gin-auth-service/src/app/services"
	"github.com/gin-gonic/gin"
)

func SetMiddlewareJSON() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("Content-Type", "application/json")
		ctx.Next()
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := services.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})

		}

		ctx.Next()

	}
}
