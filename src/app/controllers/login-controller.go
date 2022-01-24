package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	"github.com/bytesfield/golang-gin-auth-service/src/app/responses"
	"github.com/bytesfield/golang-gin-auth-service/src/app/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(c *gin.Context) {
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

	err = user.Validate("login")

	if err != nil {
		responses.ValidationError(c, "Validation Error", err)
		return
	}

	token, err := server.SignIn(user.Email, user.Password)

	if err != nil {
		//formattedError := formaterror.FormatError(err.Error())
		responses.ValidationError(c, "Validation Error", err)
		return
	}

	userData := map[string]interface{}{"token": token, "user": user}

	responses.Ok(c, "Login successfully", userData)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Model(models.User{}).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return services.CreateToken(uint32(user.ID))
}
