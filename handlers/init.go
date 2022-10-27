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
	// AUTH HANDLER
	AuthenticateUser(credentials datatransfers.UserLogin) (token string, err error)
	RegisterUser(credentials datatransfers.UserSignup) (err error)

	// USER HANDLER
	RetrieveUser(id uint) (user models.User, err error)
	UpdateUser(id uint, user datatransfers.UserUpdate) (err error)

	// TODO HANDLER
	GetAllUserTodo(user_id uint) (todo []models.Todo, err error)
	GETTodoByID(user_id uint, id uint, isAdmin bool) (todo models.Todo, err error)
	UpdateTodoUser(user_id uint, id uint, isAdmin bool, todo datatransfers.TodoUpdate) (err error)
	InsertTodo(user_id uint, newTodo datatransfers.TodoInsert) (todo models.Todo, err error)
	DeleteTodoById(user_id uint, id uint, isAdmin bool) (todo models.Todo, err error)
}

type module struct {
	db *dbEntity
}

type dbEntity struct {
	conn           *gorm.DB
	userRepository models.UserRepository
	todoRepository models.TodoRepository
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
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
			todoRepository: models.NewTodoRepository(db),
		},
	}

	return
}
