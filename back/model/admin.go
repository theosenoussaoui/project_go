package model

import (
	"encoding/json"
	"errors"
	"time"
	"regexp"
)

type Admin struct {
	User
	IsAdmin	bool
}

func (a Admin) CreateUser() *User {
	
}