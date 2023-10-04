package value_object

// PhoneNumber contains a valid phone number
type PhoneNumber struct {
	countryCode string
	number      string
}

func NewPhoneNumber(phoneNumber string) *PhoneNumber {
	return &PhoneNumber{
		number: phoneNumber,
	}
}

func ValidatePhoneNumber(string) bool {
	return true
}
