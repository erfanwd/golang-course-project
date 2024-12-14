package common

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"unicode"

	"github.com/erfanwd/golang-course-project/config"
)

// var (
// 	lowerCharSet = "abcdefghijklmnopqrstuvwxyz"
// 	upperCharSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
// 	specialCharSet = "!@#$%&*"
// 	numberSet = "0123456789"
// 	allCharSet = lowerCharSet + upperCharSet + specialCharSet + numberSet
// )

// func generatePassword () string {

// }

func PasswordValidate(value string) bool {
	cfg := config.GetConfig()
	if len(value) < cfg.Password.MinLength {
		return false
	}
	if cfg.Password.IncludeChars && !HasLetters(value) {
		return false
	}
	if cfg.Password.IncludeDigits && !HasDigits(value) {
		return false
	}
	if cfg.Password.IncludeLowercase && !HasLower(value) {
		return false
	}
	if cfg.Password.IncludeUppercase && !HasUpper(value) {
		return false
	}
	return true
}

func HasLetters(value string) bool {
	for _, item := range value {
		if unicode.IsLetter(item) {
			return true
		}
	}
	return false
}

func HasDigits(value string) bool {
	for _, item := range value {
		if unicode.IsDigit(item) {
			return true
		}
	}
	return false
}

func HasLower(value string) bool {
	for _, item := range value {
		if unicode.IsLetter(item) && unicode.IsLower(item) {
			return true
		}
	}
	return false
}

func HasUpper(value string) bool {
	for _, item := range value {
		if unicode.IsLetter(item) && unicode.IsUpper(item) {
			return true
		}
	}
	return false
}

func GenerateOtp() string {
	cfg := config.GetConfig()
	digits := cfg.Otp.Digits

	otp := ""
	for i := 0; i < digits; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(10)) 
		
		otp += fmt.Sprintf("%d", num)  
	}
	return otp
}
