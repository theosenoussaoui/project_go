package db

import (
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type PostGresSQL struct {
	db *gorm.DB
}

func (p PostGresSQL) Connect() (Persist, error) {
	// login := "host=35.187.178.167 port=5432 user=henri dbname=postgres password=admin"
	login := "host=35.187.178.167 port=" + os.Getenv("DB_PORT") + " user=" + os.Getenv("DB_USER") + " dbname=" + os.Getenv("DB_NAME") + " password=" + os.Getenv("DB_PWRD")
	db, err := gorm.Open("postgres", login)
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	db.AutoMigrate(&User{}, &Proposal{})
	p.db = db
	return &p, nil
}

func (p PostGresSQL) SaveUser(u User) error {
	return p.db.Create(&u).Error
}

func (p PostGresSQL) GetUser() ([]User, error) {
	var us []User
	err := p.db.Find(&us).Error
	return us, err
}

func (p PostGresSQL) SaveProposal(pr Proposal) error {
	return p.db.Create(&pr).Error
}

func (p PostGresSQL) GetProposal() ([]Proposal, error) {
	var ps []Proposal
	err := p.db.Find(&ps).Error
	return ps, err
}