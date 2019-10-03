package models

import (
    // "encoding/json"
    "time"
    "github.com/satori/go.uuid"
    "github.com/jinzhu/gorm"
)

type Vote struct {
    Id        	int    		`json:id`
    Uuid		string		`json:"uuid"`
    Title 		string 		`json:"title"`
    Description string 		`json:"description"`
    UUIDVote  	[]string 	`json:"uuid_vote"`
    StartDate 	time.Time 	`json:"start_date"`
    EndDate 	time.Time 	`json:"end_date"`
    IsDeleted	bool		`json:"is_boolean"`
    CreatedAt 	time.Time 	`json:"created_at"`
    UpdatedAt 	time.Time 	`json:"updated_at"`
    DeletedAt 	time.Time 	`json:"deleted_at"`
}

// GORM hook that's used for setting column created with the current Datetime 
// before any creation. Also created a new Uuid.

func (vote *Vote) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedAt", time.Now())
    scope.SetColumn("Uuid", uuid.NewV4())
    return nil
}

// GORM hook that's used for setting column created with the updated Datetime
// before any update 

func (vote *Vote) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("UpdatedAt", time.Now())
    return nil
}

// GORM hook that's used for setting column DeletedAt with the updated Datetime
// before any deletion. Also sets sets the column
// isDeleted to true.

func (vote *Vote) BeforeDelete(scope *gorm.Scope) error {
    scope.SetColumn("DeletedAt", time.Now())
    scope.SetColumn("IsDeleted", true)
    return nil
}