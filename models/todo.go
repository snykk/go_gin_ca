package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `json:"user_id"`
	Activity  string    `json:"activity"`
	Priority  string    `json:"priority"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type TodoRepository interface {
	GetAllUserTodo(user_id uint) (todos []Todo, err error)
	GetTodoByID(id uint) (todo Todo, err error)
	InsertTodo(todo Todo) (Todo, error)
	UpdateTodo(todo Todo) error
	DeleteTodo(id uint) (err error)
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) GetAllUserTodo(user_id uint) (todos []Todo, err error) {
	result := r.db.Model(&Todo{}).Where("user_id = ?", user_id).Find(&todos)
	return todos, result.Error
}

func (r *todoRepository) GetTodoByID(id uint) (todo Todo, err error) {
	result := r.db.Model(&Todo{}).Where("id = ?", id).First(&todo)
	return todo, result.Error
}

func (r *todoRepository) InsertTodo(todo Todo) (Todo, error) {
	result := r.db.Model(&Todo{}).Create(&todo)
	return todo, result.Error
}

func (r *todoRepository) UpdateTodo(todo Todo) error {
	result := r.db.Model(&Todo{}).Model(&todo).Updates(&todo)
	fmt.Println(result.Error)
	return result.Error
}

func (r *todoRepository) DeleteTodo(id uint) (err error) {
	err = r.db.Delete(&Todo{}, id).Error
	return
}
