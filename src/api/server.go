package api

import (
	"fmt"
	"log"
	"os"

	"github.com/bytesfield/golang-gin-auth-service/src/api/controllers"
	"github.com/bytesfield/golang-gin-auth-service/src/api/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error = godotenv.Load()

	if err != nil {
		log.Fatalf("An Error occured when getting env %v", err)
	} else {
		fmt.Println("Env values accessed successfully")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	server.Run(":8080")

}
