package repositories

import (
	"log"
	"time"

	"github.com/bytesfield/golang-gin-auth-service/src/app/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	SaveUser(db *gorm.DB) (*models.User, error)
	FindAllUsers(db *gorm.DB) ([]models.User, error)
	FindUserByID(db *gorm.DB, uid uint32) (*models.User, error)
	UpdateAUser(db *gorm.DB, uid uint32) (*models.User, error)
	DeleteAUser(db *gorm.DB, uid uint32) (int64, error)
}

type userRepository struct {
	user *models.User
}

func New(user *models.User) *userRepository {
	return &userRepository{
		user: user,
	}
}

func (userRepo *userRepository) SaveUser(db *gorm.DB) (*models.User, error) {
	err := userRepo.user.BeforeSave(db)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Create(&userRepo.user).Error

	if err != nil {
		return userRepo.user, err
	}

	return userRepo.user, nil

}

func (userRepo *userRepository) FindAllUsers(db *gorm.DB) (*models.User, error) {
	var err error

	users := &userRepo.user

	err = db.Find(users).Error

	if err != nil {
		return userRepo.user, err
	}
	return *users, err

}

func (userRepo *userRepository) FindUserByID(db *gorm.DB, uid uint32) (*models.User, error) {
	var err error = db.Where("id = ?", uid).First(&userRepo.user).Error

	if err != nil {
		return userRepo.user, err
	}

	return userRepo.user, err

}

func (userRepo *userRepository) UpdateUser(db *gorm.DB, uid uint32) (*models.User, error) {
	err := userRepo.user.BeforeSave(db)

	if err != nil {
		log.Fatal(err)
	}

	db = db.Model(&userRepo.user).Where("id = ?", uid).Updates(
		map[string]interface{}{
			"password":   userRepo.user.Password,
			"firstname":  userRepo.user.Firstname,
			"lastname":   userRepo.user.Lastname,
			"email":      userRepo.user.Email,
			"updated_at": time.Now(),
		},
	)

	if db.Error != nil {
		return userRepo.user, db.Error
	}

	err = db.Debug().Model(&userRepo.user).Where("id = ?", uid).Take(&userRepo.user).Error

	if err != nil {
		return userRepo.user, err
	}

	return userRepo.user, nil

}

func (userRepo *userRepository) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {

	err := db.Delete(&userRepo.user, uid).Error

	if err != nil {
		return 0, err
	}
	return 1, nil

}
