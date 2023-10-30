package entity

import (
	"github.com/google/uuid"
	vo "github.com/moswil/course-management/user-service/internal/core/domain/value_object"
)

// UserEntity
type UserEntity struct {
	userId        uuid.UUID
	email         vo.Email
	username      string
	phoneNumber   vo.PhoneNumber
	firstName     string
	middleName    string
	lastName      string
	password      string
	emailVerified bool
}

func NewUserEntity(email, username, phoneNumber, firstName, middleName, lastName, password string) *UserEntity {
	return &UserEntity{
		userId:      uuid.New(),
		email:       *vo.NewEmail(email),
		username:    username,
		phoneNumber: *vo.NewPhoneNumber(phoneNumber),
		firstName:   firstName,
		middleName:  middleName,
		lastName:    lastName,
		password:    password,
	}
}
