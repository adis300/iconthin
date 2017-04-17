package main

import (
	"regexp"
)

var emailRegex = regexp.MustCompile(`^([\w\.\_]{1,32})@(\w{2,32})\.(\w{1,20}\.)?([\w]{2,8})$`)

// validateEmail checks if an email is valid
func validateEmail(email string) bool {
	if len(email) < 6 || len(email) > 64 {
		return false
	}
	return emailRegex.Match([]byte(email))
}
