package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	UserId string `gorm:"column:user_id"`
	AuthId string `gorm:"column:auth_id"`
}

// TableName sets the table name for the UserModel in the database.
func (UserModel) TableName() string {
	return "users" // Customize the table name as needed
}
