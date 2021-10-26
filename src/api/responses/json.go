package responses

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

func Ok(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, true, message, http.StatusOK, data)
}

func CreatedResponse(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusCreated, data)
}

func BadRequest(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusBadRequest, data)
}

func ValidationError(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusUnprocessableEntity, data)
}

func ServerError(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusInternalServerError, data)
}

func ConflictResponse(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusConflict, data)
}

func ForbiddenResponse(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusForbidden, data)
}

func NotFound(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusNotFound, data)
}

func ServiceUnavailableResponse(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusServiceUnavailable, data)
}

func Unauthorized(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusUnauthorized, data)
}

func NoContent(c *gin.Context, message string, data ...interface{}) {

	BuildResponse(c, false, message, http.StatusNoContent, data)
}

func BuildResponse(c *gin.Context, status bool, message string, statusCode int, data interface{}) {

	if reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface()) {

		c.JSON(statusCode, gin.H{"message": message, "status": status})
	} else {
		c.JSON(statusCode, gin.H{"status": status, "message": message, "data": data})
	}

}
