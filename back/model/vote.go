package model

import (
	"encoding/json"
	"time"
)

type User struct {
	Id        	int    		`json:id`
	Uuid		string		`json:"uuid"`
	Title 		string 		`json:"title"`
	Description string 		`json:"description"`
	UUIDVote  	[]string 	`json:"uuid_vote"`
	StartDate 	time.Time 	`json:"start_date"`
	EndDate 	time.Time 	`json:"end_date"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	DeletedAt 	time.Time 	`json:"deleted_at"`
}

