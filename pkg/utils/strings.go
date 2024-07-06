package utils

import "unicode"

const specialCharSet string = "!@#$%&*"

func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) {
			return true
		}
	}
	return false
}

func HasSpecial(s string) bool {
	for _, i := range s {
		for _, j := range specialCharSet {
			if i == j {
				return true
			}
		}
	}
	return false
}

func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) {
			return true
		}
	}
	return false
}

func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func HasDigit(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
