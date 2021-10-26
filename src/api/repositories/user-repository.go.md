<!-- @format -->

package repositories

import (
"log"
"time"

    "github.com/bytesfield/golang-gin-auth-service/src/api/models"
    "gorm.io/gorm"

)

type UserRepository interface {
SaveUser(db *gorm.DB) (*models.User, error)
FindAllUsers(db _gorm.DB) (_[]models.User, error)
FindUserByID(db *gorm.DB, uid uint32) (*models.User, error)
UpdateAUser(db *gorm.DB, uid uint32) (*models.User, error)
DeleteAUser(db \*gorm.DB, uid uint32) (int64, error)
}

type userRepository struct {
user []models.User
}

func New() UserRepository {
return &userRepository{} //A Pointer to videoService struct that implements the VideoService Interface
}

func (u *userRepository) SaveUser(db *gorm.DB) (\*models.User, error) {

    var err error = db.Debug().Create(&u).Error
    if err != nil {
    	return &u.user{}, err
    }
    return u, nil

}

func (u *userRepository) FindAllUsers(db *gorm.DB) (\*[]models.User, error) {
var err error
users := u.user

    err = db.Debug().Model(&u.user).Limit(100).Find(&users).Error

    if err != nil {
    	return &u.user, err
    }
    return &users, err

}

func (u *userRepository) FindUserByID(db *gorm.DB, uid uint32) (\*models.User, error) {
var err error = db.Debug().Model(u.user).Where("id = ?", uid).Take(&u).Error

    if err != nil {
    	return &u.user, err
    }
    // if gorm.IsRecordNotFoundError(err) {
    // 	return &User{}, errors.New("User Not Found")
    // }
    return u, err

}

func (u *userRepository) UpdateAUser(db *gorm.DB, uid uint32) (\*models.User, error) {

    // To hash the password
    err := u.BeforeSave()
    if err != nil {
    	log.Fatal(err)
    }
    db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
    	map[string]interface{}{
    		"password":  u.Password,
    		"firstname": u.Firstname,
    		"lastname":  u.Lastname,
    		"email":     u.Email,
    		"update_at": time.Now(),
    	},
    )
    if db.Error != nil {
    	return &User{}, db.Error
    }
    // This is the display the updated user
    err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
    if err != nil {
    	return &User{}, err
    }
    return u, nil

}

func (u *userRepository) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

    db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

    if db.Error != nil {
    	return 0, db.Error
    }
    return db.RowsAffected, nil

}
