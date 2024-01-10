package main

import (
	"github.com/pquerna/otp/totp"
)

func GenerateOTPURI(label string, secretKey string) string {
	// Generate the OTP URI
	otpURI, _ := totp.GenerateCodeCustom(label, secretKey, totp.ValidateOpts{
		Period:    30,     // OTP validity period (seconds)
		Digits:    6,      // OTP length
		Algorithm: "SHA1", // Hash algorithm
	})

	return otpURI
}
