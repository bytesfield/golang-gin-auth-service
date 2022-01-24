package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Firstname string `gorm:"size:255;not null" json:"firstname"`
	Lastname  string `gorm:"size:255;not null" json:"lastname"`
	Nickname  string `gorm:"size:255;null;unique" json:"nickname"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave(*gorm.DB) error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Firstname = html.EscapeString(strings.TrimSpace(u.Firstname))
	u.Lastname = html.EscapeString(strings.TrimSpace(u.Lastname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Firstname == "" {
			return errors.New("firstname required")
		}
		if u.Lastname == "" {
			return errors.New("lastname required")
		}
		if u.Password == "" {
			return errors.New("password required")
		}
		if u.Email == "" {
			return errors.New("email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email format")
		}

		return nil
	case "login":
		if u.Password == "" {
			return errors.New("password required")
		}
		if u.Email == "" {
			return errors.New("email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email format")
		}
		return nil

	default:
		if u.Firstname == "" {
			return errors.New("firstname required")
		}
		if u.Lastname == "" {
			return errors.New("lastname required")
		}
		if u.Password == "" {
			return errors.New("password required")
		}
		if u.Email == "" {
			return errors.New("email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email format")
		}
		return nil
	}
}
