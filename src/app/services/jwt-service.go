package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateToken(user_id uint32) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func VerifyToken(ctx *gin.Context) error {
	tokenString, err := GetToken(ctx)

	if err != nil {
		return err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			err := errors.New("invalid Signature")

			ctx.AbortWithStatus(http.StatusBadRequest)

			return err
		}
		return err
	}

	// Check if the token is valid.
	if !token.Valid {
		err := errors.New("the token is not valid")

		ctx.AbortWithStatus(http.StatusBadRequest)

		return err
	}

	//Extract key value from the token and print them on console
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		Pretty(claims)
	}

	return nil
}

func GetToken(ctx *gin.Context) (string, error) {
	authorizationHeader := ctx.GetHeader("Authorization")

	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")

		ctx.AbortWithStatus(http.StatusBadRequest)

		return "", err
	}

	fields := strings.Fields(authorizationHeader)

	if len(fields) < 2 {

		err := errors.New("invalid Authorization header format")

		ctx.AbortWithStatus(http.StatusBadRequest)

		return "", err
	}

	authorizationType := strings.ToLower(fields[0])

	if strings.ToLower(authorizationType) != "bearer" {

		err := errors.New("unsupported authorization type")

		ctx.AbortWithStatus(http.StatusBadRequest)

		return "", err

	}

	accessToken := fields[1]

	return accessToken, nil
}

func GetTokenID(ctx *gin.Context) (uint32, error) {

	tokenString, err := GetToken(ctx)

	if err != nil {
		return 0, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)

		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}

	return 0, nil
}

//Pretty display the claims nicely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
