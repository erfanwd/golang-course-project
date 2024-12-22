package service_errors

const (
	OtpExists     string = "otp exists"
	OtpUsed       string = "otp used"
	OtpIsNotValid string = "otp is not valid"

	UnexpectedError string = "unexpected error"
	ClaimsNotFound  string = "claims not found"

	UsernameExists     string = "username exists"
	UsernameNotExists  string = "username not exists"
	EmailExists        string = "email exists"
	MobileNumberExists string = "mobile number exists"
	IncorrectPassword  string = "password is incorrect"

	// Token
	UnExpectedError = "Expected error"
	TokenRequired   = "token required"
	TokenExpired    = "token expired"
	TokenInvalid    = "token invalid"
)
