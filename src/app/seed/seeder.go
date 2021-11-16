package seed

import (
	"fmt"
	"log"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	"github.com/jaswdr/faker"
	"gorm.io/gorm"
)

func Load(db *gorm.DB) {

	faker := faker.New()

	hashedPassword, err := models.Hash("password")

	if err != nil {
		fmt.Println(err)
	}

	password := string(hashedPassword)

	var users = []models.User{
		models.User{
			Firstname: faker.Person().FirstName(),
			Lastname:  faker.Person().LastName(),
			Email:     faker.Person().Contact().Email,
			Nickname:  faker.Person().FirstName(),
			Password:  password,
		},
		models.User{
			Firstname: faker.Person().FirstName(),
			Lastname:  faker.Person().LastName(),
			Email:     faker.Person().Contact().Email,
			Nickname:  faker.Person().FirstName(),
			Password:  password,
		},
	}

	err = db.Debug().Migrator().DropTable(&models.User{})

	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}

	err = db.Debug().AutoMigrate(&models.User{})

	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error

		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

}
