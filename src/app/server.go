package app

import (
	"log"
	"os"

	"github.com/bytesfield/golang-gin-auth-service/src/app/controllers"
	"github.com/bytesfield/golang-gin-auth-service/src/app/seed"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("An Error occurred when getting .env file %v", err)
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	seed.Load(server.DB)

	//server.Router.LoadHTMLGlob("../api/templates/*.html")

	server.Run(":8080")

}
