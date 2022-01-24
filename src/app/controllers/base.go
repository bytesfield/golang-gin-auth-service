package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	c "github.com/bytesfield/golang-gin-auth-service/src/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/bytesfield/golang-gin-auth-service/src/app/middlewares"
	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	gin "github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func initializeConfig() {

	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)

	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	fmt.Println("Configurations read successfully")

}

func (server *Server) Initialize(DBdriver string, DBUser string, DBPassword, DBPort, DBHost, DBName string) {
	var err error

	initializeConfig()

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

	setUpLogOutput()

	server.DB.AutoMigrate(&models.User{}) //database migration

	server.Router = gin.New()

	server.Router.Use(gin.Recovery(), middlewares.Logger())

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")

	log.Fatal(http.ListenAndServe(addr, server.Router))
}
