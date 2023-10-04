package dto

// CreateUserDTO payload to create a user
type CreateUserDTO struct {
	Email       string `json:"email"`
	Username    string `json:"username"`
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name,omitempty"`
	LastName    string `json:"last_name"`
	Password    string `json:"password"`
}

// UserDTO for retrieving a user
type UserDTO struct {
	UserId string `json:"user_id"`
}
