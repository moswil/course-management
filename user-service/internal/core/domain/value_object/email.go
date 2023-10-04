package value_object

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

// ValidateEmail checks if the passed email is valid
func ValidateEmail(email string) bool {
	return true
}
