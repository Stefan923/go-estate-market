package util

import "unicode"

func HasUpper(value string) bool {
	for _, character := range value {
		if unicode.IsUpper(character) && unicode.IsLetter(character) {
			return true
		}
	}
	return false
}

func HasLower(value string) bool {
	for _, character := range value {
		if unicode.IsLower(character) && unicode.IsLetter(character) {
			return true
		}
	}
	return false
}

func HasLetter(value string) bool {
	for _, character := range value {
		if unicode.IsLetter(character) {
			return true
		}
	}
	return false
}

func HasDigits(value string) bool {
	for _, character := range value {
		if unicode.IsDigit(character) {
			return true
		}
	}
	return false
}
