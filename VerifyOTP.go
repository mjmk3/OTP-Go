package main

import "github.com/pquerna/otp/totp"

func VerifyOTP(otp string, secretKey string) bool {
	// Verify the OTP
	valid := totp.Validate(otp, secretKey)

	return valid
}
