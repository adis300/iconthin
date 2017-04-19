package main

import (
	"regexp"
	"strings"
)

// 33! 35# 36$ 37% 38& 42* 43+ 44, 45- 46. 58: 63? 64@ 94^ 95_ 126~
// Only allows the following chars !#$%&*+,-.:?@^_~
var passwordSpecialChars = []byte{33, 35, 36, 37, 38, 42, 43, 44, 45, 46, 58, 63, 64, 94, 95, 126}
var emailRegex = regexp.MustCompile(`^([\w\.\_]{1,32})@(\w{2,32})\.(\w{1,20}\.)?([\w]{2,8})$`)

// validateEmail checks if an email is valid
func validateEmail(email string) bool {
	if len(email) < 6 || len(email) > 64 {
		return false
	}
	return emailRegex.Match([]byte(email))
}

func validateToken(token string) bool {
	return token == currentToken
}

// validateUsername validates username and eliminate all bad chars
func validateUsername(username string) bool {
	if len(username) > 32 || len(username) < 4 {
		return false
	}
	raw := []byte(strings.ToLower(username))
	for i := range raw {
		// Eliminating invalid characters from ascii table
		if raw[i] < 48 || (raw[i] > 57 && raw[i] < 97) || raw[i] > 122 {
			return false
		}
	}
	return true
}

// validatePassword validates password and eliminate all bad chars
func validatePassword(password string) bool {
	if len(password) > 32 || len(password) < 6 {
		return false
	}
	raw := []byte(strings.ToLower(password))
	for i := range raw {
		// Eliminating invalid characters from ascii table
		if raw[i] < 48 || (raw[i] > 57 && raw[i] < 97) || raw[i] > 122 {
			for j := range passwordSpecialChars {
				if raw[i] == passwordSpecialChars[j] {
					// The character is one of the valid special password chars
					continue
				}
			}
			return false
		}
	}
	return true
}
