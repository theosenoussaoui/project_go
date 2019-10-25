package db

// Import
import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/theosenoussaoui/project_go/back/models"
	"github.com/qor/validations"
)

var db *gorm.DB
var err error

// Initialize creates a connection to postgres database and 
// migrates all our models
func Initialize() {
	err := godotenv.Load()

	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	// GoDotEnv get our evironment variables from our .env file.

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PWRD")

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	// Opening database connection with Gorm.

	db, err = gorm.Open("postgres", dbinfo)
	db.LogMode(true)

	validations.RegisterCallbacks(db)

	if err != nil {
		log.Println("Failed to connect to database")
		panic(err)
	}

	log.Println("Database connected")

	// Creating User table if it doesn't exist already.

	if !db.HasTable(&models.User{}) {
		err := db.CreateTable(&models.User{})
		if err != nil {
			log.Println("Table User already exists.")
		} else {
			log.Println("Table User was successfully created.")
		}
	}

		// Vote Table
		if !db.HasTable(&models.Vote{}) {
			err := db.CreateTable(&models.Vote{})
			if err != nil {
				log.Println("Table Votes already exists.")
			} else {
				log.Println("Table Votes was successfully created.")
			}
		}
	
		// Blacklist Table
		if !db.HasTable(&models.Blacklist{}) {
			err := db.CreateTable(&models.Blacklist{})
			if err != nil {
				log.Println("Table Blacklist already exists.")
			} else {
				log.Println("Table Blacklist was successfully created.")
			}
		}

		// Run migrations (models and field changes)
		db.AutoMigrate(&models.User{})
		db.AutoMigrate(&models.Vote{})
		db.AutoMigrate(&models.Blacklist{})
}

// Creating administrator based on the User model 

func CreateSystemAdmin() {
	user := models.User{
		FirstName:   "test",
		LastName:    "test",
		AccessLevel: 1,
		Email:       "test@gmail.com",
	}

	user.SetPassword("test")

	dateOfBirth, err := time.Parse("02-01-2006", "01-01-1990")
	if err != nil {
		log.Fatal(err)
	}
	user.DateOfBirth = dateOfBirth

	db.NewRecord(user) // => returns `true` as primary key is blank

	if err := db.Create(&user); err.Error != nil {
		log.Println("Error creating CreateSystemAdmin : " + user.FirstName + " " + user.LastName)
		log.Println(err)
		return
	}
}

// GetDB get gorm db instance

func GetDB() *gorm.DB {
	return db
}

// CloseDB close gorm db instance

func CloseDB() {
	db.Close()
}
