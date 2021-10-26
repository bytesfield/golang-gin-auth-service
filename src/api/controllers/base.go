package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bytesfield/golang-gin-auth-service/src/api/middlewares"
	"github.com/bytesfield/golang-gin-auth-service/src/api/models"
	gin "github.com/gin-gonic/gin"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

//Writing logs to file
func setUpLogOutput() {

	f, err := os.Create("gin.log")

	if err != nil {
		println("Could not create file")
	}

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func (server *Server) Initialize(DBdriver string, DBUser string, DBPassword, DBPort, DBHost, DBName string) {

	setUpLogOutput()

	var err error

	if DBdriver == "mysql" {
		DBUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

		server.DB, err = gorm.Open(mysql.Open(DBUrl), &gorm.Config{})

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DBdriver)

			log.Fatal("Error:", err)
		} else {
			fmt.Printf("Connected to %s database successfully", DBdriver)
		}
	}

	if DBdriver == "postgres" {

		DBUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)

		server.DB, err = gorm.Open(postgres.Open(DBUrl), &gorm.Config{})

		if err != nil {
			fmt.Printf("Cannot connect to %s database", DBdriver)

			log.Fatal("Error:", err)
		} else {
			fmt.Printf("Connected to %s database successfully", DBdriver)
		}
	}

	server.DB.Debug().AutoMigrate(&models.User{}) //database migration

	server.Router = gin.New()

	server.Router.Use(gin.Recovery(), middlewares.Logger())

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
