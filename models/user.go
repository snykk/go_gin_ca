package models

import (
	"time"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex" json:"username"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserRepository interface {
	GetUserByID(id uint) (user User, err error)
	GetUserByUsername(username string) (user User, err error)
	InsertUser(user User) (id uint, err error)
	UpdateUser(user User) (err error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(id uint) (user User, err error) {
	result := r.db.Model(&User{}).Where("id = ?", id).First(&user)
	return user, result.Error
}

func (r *userRepository) GetUserByUsername(username string) (user User, err error) {
	result := r.db.Model(&User{}).Where("username = ?", username).First(&user)
	return user, result.Error
}

func (r *userRepository) InsertUser(user User) (id uint, err error) {
	result := r.db.Model(&User{}).Create(&user)
	return user.ID, result.Error
}

func (r *userRepository) UpdateUser(user User) (err error) {
	result := r.db.Model(&User{}).Model(&user).Updates(&user)
	return result.Error
}
