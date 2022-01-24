package testcase

import (
	"log"
	"testing"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	userRepository "github.com/bytesfield/golang-gin-auth-service/src/app/repositories"
	"github.com/jaswdr/faker"
	"gopkg.in/go-playground/assert.v1"
	_ "gorm.io/driver/mysql"
	_ "gorm.io/driver/postgres"
)

func TestFindAllUsers(t *testing.T) {

	refreshUserTable()

	seededUsers, err := seedUsers()

	if err != nil {
		log.Fatal(err)
	}

	userRepo := userRepository.New(&userInstance)

	users, err := userRepo.FindAllUsers(server.DB)

	if err != nil {
		t.Errorf("error getting the users: %v\n", err)
		return
	}

	assert.IsEqual(seededUsers, *users)
}

func TestSaveUser(t *testing.T) {

	refreshUserTable()

	faker := faker.New()

	newUser := models.User{
		Firstname: faker.Person().FirstName(),
		Lastname:  faker.Person().LastName(),
		Email:     faker.Person().Contact().Email,
		Nickname:  faker.Person().FirstName(),
		Password:  "password",
	}

	userRepo := userRepository.New(&userInstance)

	savedUser, err := userRepo.SaveUser(server.DB)

	if err != nil {
		t.Errorf("error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nickname, savedUser.Nickname)
}

func TestGetUserByID(t *testing.T) {

	refreshUserTable()

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	userRepo := userRepository.New(&userInstance)

	foundUser, err := userRepo.FindUserByID(server.DB, uint32(user.ID))

	if err != nil {
		t.Errorf("error getting one user: %v\n", err)
		return
	}

	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)
}

func TestUpdateAUser(t *testing.T) {

	refreshUserTable()

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	faker := faker.New()

	userUpdate := models.User{
		Firstname: faker.Person().FirstName(),
		Lastname:  faker.Person().LastName(),
		Email:     faker.Person().Contact().Email,
		Nickname:  faker.Person().FirstName(),
		Password:  "password",
	}
	userRepo := userRepository.New(&userInstance)

	updatedUser, err := userRepo.UpdateUser(server.DB, uint32(user.ID))

	if err != nil {
		t.Errorf("error updating the user: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Nickname, userUpdate.Nickname)
}

func TestDeleteAUser(t *testing.T) {

	refreshUserTable()

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	userRepo := userRepository.New(&userInstance)

	isDeleted, err := userRepo.DeleteUser(server.DB, uint32(user.ID))

	if err != nil {
		t.Errorf("error updating the user: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}
