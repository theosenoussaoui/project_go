package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"encoding/json"
	"errors"
	"time"
	"regexp"
)

type User struct {
	gorm.Model
	Id        	int    		`json:id`
	FirstName 	string 		`json:"first_name"`
	LastName  	string 		`json:"last_name"`
	Email  		string 		`json:"email"`
	DateOfBirth time.Time 	`json:"date_of_birth"`
	Password  	string 		`json:"pass"`
}

func (u User) Valid() []error {
	
	var errs []error
	regexNumbers 	:= regexp.MustCompile("[0-9]+")
	regexSpace		:= regexp.MustCompile("\s+")
	regexEmail		:= regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	
	if len(u.FirstName) < 2 {
		errs = append(errs, errors.New("First name field must be at least 2 characters long."))
	}

	if regexNumbers.MatchString(u.FirstName) {
		errs = append(errs, errors.New("First name field contains number(s)."))
	}

	if regexSpace.MatchString(u.FirstName) {
		errs = append(errs, errors.New("First name field contains space(s)."))
	}

	if len(u.LastName) < 2 {
		errs = append(errs, errors.New("Last name field must be at least 2 characters long."))
	}

	if regexNumbers.MatchString(u.LastName) {
		errs = append(errs, errors.New("Last name field contains numbers."))
	}

	if regexSpace.MatchString(u.LastName) {
		errs = append(errs, errors.New("Last name field contains space(s)."))
	}

	if regexEmail.MatchString(u.Email) {
		errs = append(errs, errors.New("Email is invalid."))
	}

	if u.DateOfBirth.IsZero() {
		errs = append(errs, errors.New("Date of birth field is empty."))
	}

	// vérifier que le user soit majeur
	
	if len(u.Password) == 0 {
		errs = append(errs, errors.New("Password field is empty."))
	}
	
	if len(errs) != 0 {
		return errs
	}
	
	return nil
}