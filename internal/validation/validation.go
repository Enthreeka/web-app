package validation

import "regexp"

var (
	passwordPattern = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()_+-=.,:;'"]{6,}$`)
	loginPattern    = regexp.MustCompile(`^[a-zA-Z0-9]{3,25}$`)
	namePattern     = regexp.MustCompile(`^[a-zA-Za-яА-Я]{0,10}$`)
)

func IsValidationPassword(password string) bool {
	return passwordPattern.MatchString(password)
}

func IsValidationLogin(login string) bool {
	return loginPattern.MatchString(login)
}

func IsValidationName(name string) bool {
	return namePattern.MatchString(name)
}
