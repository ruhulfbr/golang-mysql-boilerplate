package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"time"

	_ "golang.org/x/crypto/bcrypt"
)

type UserStatus int

const (
	StatusActive   UserStatus = 1 // Active
	StatusInactive UserStatus = 0 // Inactive
)

type User struct {
	Id        uint32     `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:255;not null" json:"name"`
	Email     string     `gorm:"size:255;not null;unique" json:"email"`
	Username  string     `gorm:"size:100;not null;unique" json:"username"`
	Password  string     `gorm:"size:100;not null" json:"password"`
	Status    UserStatus `gorm:"size:1;not null" json:"status"`
	LastLogin time.Time  `gorm:"type:datetime;default:null" json:"last_login"`
	CreatedAt time.Time  `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) Prepare() {
	u.Id = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
	u.LastLogin = time.Now()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Username == "" {
			return errors.New("username is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}

		return nil

	default:
		if u.Name == "" {
			return errors.New("name is required")
		}
		if u.Username == "" {
			return errors.New("username is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}

		return nil
	}
}
