package testcase

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/bytesfield/golang-gin-auth-service/src/app/controllers"
	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	"github.com/jaswdr/faker"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	var err error

	err = godotenv.Load(os.ExpandEnv("../../.env"))

	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}

	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

	if TestDbDriver == "mysql" {

		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TestDbUser"), os.Getenv("TestDbPassword"), os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbName"))

		server.DB, err = gorm.Open(mysql.Open(DBURL), &gorm.Config{})

		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to %s database\n", TestDbDriver)
		}
	}

	if TestDbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("TestDbHost"), os.Getenv("TestDbPort"), os.Getenv("TestDbUser"), os.Getenv("TestDbName"), os.Getenv("TestDbPassword"))

		server.DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})

		if err != nil {
			fmt.Printf("Cannot connect to %s database\n", TestDbDriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("Connected to the %s database\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	server.DB.Debug().Migrator().DropTable(&models.User{})

	server.DB.Debug().AutoMigrate(&models.User{})

	log.Printf("Successfully refreshed table")

	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	faker := faker.New()

	user := models.User{
		Firstname: faker.Person().FirstName(),
		Lastname:  faker.Person().LastName(),
		Email:     faker.Person().Contact().Email,
		Nickname:  faker.Person().FirstName(),
		Password:  faker.Hash().SHA256(),
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error

	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}

	return user, nil
}

func seedUsers() ([]models.User, error) {

	refreshUserTable()

	faker := faker.New()

	users := []models.User{
		models.User{
			Firstname: faker.Person().FirstName(),
			Lastname:  faker.Person().LastName(),
			Email:     faker.Person().Contact().Email,
			Nickname:  faker.Person().FirstName(),
			Password:  faker.Hash().SHA256(),
		},
		models.User{
			Firstname: faker.Person().FirstName(),
			Lastname:  faker.Person().LastName(),
			Email:     faker.Person().Contact().Email,
			Nickname:  faker.Person().FirstName(),
			Password:  faker.Hash().SHA256(),
		},
	}
	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			return []models.User{}, err
		}
	}

	return users, nil
}
