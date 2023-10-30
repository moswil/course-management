package value_object

import "regexp"

// Email validates an email
type Email struct {
	email string
}

// NewEmail creates a new email address
func NewEmail(email string) *Email {
	return &Email{
		email: email,
	}
}

// IsEmailValid returns true if the email passwed is valid, otherwise returns false.
func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
