package models

import (
	// "encoding/json"
	// "strconv"
	// "strings"
	"errors"
	"log"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	"github.com/theosenoussaoui/project_go/back/utils"
	uuid "github.com/satori/go.uuid"
)

// User structure represents the user.

type User struct {
	ID          int       `gorm:"primary_key"`
	UUID        uuid.UUID `json:"uuid"`
	AccessLevel int       `json:"access_level" valid:"range(0|1),numeric"`
	FirstName   string    `json:"first_name" valid:"required,alpha,length(2|255)"`
	LastName    string    `json:"last_name" valid:"required,alpha,length(2|255)"`
	Email       string    `json:"email" valid:"email,required"`
	Password    string    `json:"pass" valid:"required"`
	DateOfBirth time.Time `json:"birth_date" valid:"required"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// UserResponse represents user data that can be returned 
// as response to an GET/POST request on /users

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
		db.AddError(errors.New("User: age need to be 18+"))
	}

	// check if user already exists
	var u User
	if !db.Where("email = ?", user.Email).First(&u).RecordNotFound() {
		db.AddError(errors.New("User: already exists"))
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

// BeforeUpdate is gorm hook that is triggered on every updated on user struct
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}

// MarshalJSON is marshaling the user.
// Function isn't finished

// func (user User) MarshalJSON() ([]byte, error) {

// 	var ur UserResponse
// 	layout := "02-01-2006"

// 	// Custom date in dd-mm-yyy format
// 	year, month, day := user.DateOfBirth.Date()
// 	date := strconv.Itoa(day) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(year)
// 	t, _ := time.Parse(layout, date)

// 	ur.UUID = user.UUID
// 	ur.FirstName = user.FirstName
// 	ur.LastName = user.LastName
// 	ur.Email = user.Email
// 	ur.DateOfBirth = t.Format(layout)
// 	return json.Marshal(ur)
// }

// UnmarshalJSON create user formatted representation of jsonData

// func (user *User) UnmarshalJSON(data []byte) error {
// 	var jsonData map[string]string
// 	err := json.Unmarshal(data, &jsonData)
// 	if err != nil {
// 		return err
// 	}

// 	for key, value := range jsonData {
// 		if strings.ToLower(key) == "first_name" {
// 			user.FirstName = value
// 		}

// 		if strings.ToLower(key) == "last_name" {
// 			user.LastName = value
// 		}

// 		if strings.ToLower(key) == "email" {
// 			user.Email = value
// 		}

// 		if strings.ToLower(key) == "pass" {
// 			user.SetPassword(value)
// 		}
// 	}
// 	return nil
// }
