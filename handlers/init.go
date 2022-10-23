package handlers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/snykk/go_gin_ca/config"
	"github.com/snykk/go_gin_ca/datatransfers"
	"github.com/snykk/go_gin_ca/models"
)

var Handler HandlerFunc

type HandlerFunc interface {
	AuthenticateUser(credentials datatransfers.UserLogin) (token string, err error)
	RegisterUser(credentials datatransfers.UserSignup) (err error)
	RetrieveUser(id uint) (user models.User, err error)
	UpdateUser(id uint, user datatransfers.UserUpdate) (err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn           *gorm.DB
	userRepository models.UserRepository
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func InitializeHandler() (err error) {
	// Initialize DB
	var db *gorm.DB
	db, err = gorm.Open(postgres.Open(
		fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
			config.AppConfig.DBHost, config.AppConfig.DBPort, config.AppConfig.DBDatabase,
			config.AppConfig.DBUsername, config.AppConfig.DBPassword),
	), &gorm.Config{})

	if err != nil {
		log.Println("[INIT] failed connecting to PostgreSQL")
		return
	}
	log.Println("[INIT] connected to PostgreSQL")
	dbMigrate(db)

	// Compose handler modules
	Handler = &module{
		db: &dbEntity{
			conn:           db,
			userRepository: models.NewUserRepository(db),
		},
	}

	return
}
