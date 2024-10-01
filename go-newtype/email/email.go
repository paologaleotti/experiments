package email

type Email struct {
	value string
}

func EmailFromString(value string) (Email, error) {
	// VALIDATE
	return Email{value: value}, nil
}

func (e Email) String() string {
	return e.value
}
