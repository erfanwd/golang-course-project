package common

import (
	"crypto/rand"
	"fmt"
	"math/big"
	math "math/rand"
	"regexp"
	"strings"

	"unicode"

	"github.com/erfanwd/golang-course-project/config"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

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

func GeneratePassword() string {
	var password strings.Builder

	cfg := config.GetConfig()
	passwordLength := cfg.Password.MinLength + 2
	minSpecialChar := 2
	minNum := 3
	if !cfg.Password.IncludeDigits {
		minNum = 0
	}

	minUpperCase := 3
	if !cfg.Password.IncludeUppercase {
		minUpperCase = 0
	}

	minLowerCase := 3
	if !cfg.Password.IncludeLowercase {
		minLowerCase = 0
	}

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := math.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := math.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := math.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	//Set lowercase
	for i := 0; i < minLowerCase; i++ {
		random := math.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := math.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	math.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}
