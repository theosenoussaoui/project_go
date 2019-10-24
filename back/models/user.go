package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/jinzhu/gorm"
	"github.com/theosenoussaoui/project_go/back/utils"
	uuid "github.com/satori/go.uuid"
)

// User represents the user.
type User struct {
	ID          int       `gorm:"primary_key"`
	UUID        uuid.UUID `json:"uuid"`
	AccessLevel int       `json:"access_level" valid:"range(0|1),numeric"`
	FirstName   string    `json:"first_name" valid:"required,alpha,length(2|255)"`
	LastName    string    `json:"last_name" valid:"required,alpha,length(2|255)"`
	Email       string    `json:"email" valid:"email,required" gorm:"unique;not null"`
	Password    string    `json:"pass" valid:"required"`
	DateOfBirth time.Time `json:"birth_date" valid:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Blacklists  []Blacklist `json:"blacklists"`
	Votes       []*Vote     `json:"uuid_votes" gorm:"many2many:votes_users;association_foreignkey:UUID;foreignkey:uuid"`
}

// UserResponse represents user data that can be returned as response
type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateOfBirth string    `json:"birth_date"`
}

// Validate checks that user struct is valid
func (user User) Validate(db *gorm.DB) {
	// check if user is adult
	if age := utils.Age(user.DateOfBirth); age < 18 {
		db.AddError(errors.New("user: age need to be 18+"))
	}
}

// SetPassword create a hashed password for user
func (user *User) SetPassword(password string) {
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost); err != nil {
		log.Println(err)
	} else {
		user.Password = string(hashedPassword)
	}
}

// IsAdmin Determine whether user is admin or not
func (user *User) IsAdmin() bool {
	return user.AccessLevel == 1
}

// BeforeCreate is gorm hook that is triggered before saving new user
func (user *User) BeforeCreate(scope *gorm.Scope) error {
	// or error handling
	u2 := uuid.NewV4()
	
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UUID", u2)
	return nil
}

// BeforeUpdate is gorm hook that is triggered on every updated on user struct
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// MarshalJSON is marshaling the user.
func (user User) MarshalJSON() ([]byte, error) {

	var ur UserResponse
	layout := "02-01-2006"
	t := user.DateOfBirth
	dateOfBirth := fmt.Sprintf("%s", t.Format(layout))

	ur.UUID = user.UUID
	ur.FirstName = user.FirstName
	ur.LastName = user.LastName
	ur.Email = user.Email
	ur.DateOfBirth = dateOfBirth
	return json.Marshal(ur)
}

// UnmarshalJSON create user formatted representation of jsonData
func (user *User) UnmarshalJSON(data []byte) error {
	var jsonData map[string]string
	err := json.Unmarshal(data, &jsonData)
	if err != nil {
		return err
	}

	for key, value := range jsonData {
		if strings.ToLower(key) == "first_name" {
			user.FirstName = value
		}

		if strings.ToLower(key) == "last_name" {
			user.LastName = value
		}

		if strings.ToLower(key) == "email" {
			user.Email = value
		}

		if strings.ToLower(key) == "pass" {
			user.SetPassword(value)
		}

		if strings.ToLower(key) == "access_level" {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return err
			}
			user.AccessLevel = intValue
		}

		if strings.ToLower(key) == "birth_date" {
			dateOfBirth, err := time.Parse("02-01-2006", value)
			if err != nil {
				return err
			}
			user.DateOfBirth = dateOfBirth
		}
	}
	return nil
}