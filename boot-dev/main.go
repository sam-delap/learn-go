package main

import "unicode"

func isValidPassword(password string) bool {
	passwordLen := len(password)

	if passwordLen < 5 || passwordLen > 12 {
		return false
	}

	hasUpper := false
	hasDigit := false
	for _, c := range password {

		if unicode.IsUpper(c) {
			hasUpper = true
		}

		if unicode.IsDigit(c) {
			hasDigit = true
		}
		
		if hasUpper && hasDigit {
			return true
		}
	}

	return false
}

