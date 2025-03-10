package models

import (
	"ticket-booking/utils" // Import the utils package where DB is initialized
	"github.com/jinzhu/gorm"
)

// User model
type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// Save user to DB
func (user *User) Save() error {
	return utils.DB.Create(&user).Error
}

// Get user by username
func GetUserByUsername(username string) (*User, error) {
	var user User
	if err := utils.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
