package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/generate", func(c *gin.Context) {
		// Generate a random secret key
		secretKey, err := generateRandomSecretKey()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate secret key"})
			return
		}

		// Generate OTP URI for the secret key
		otpURI := GenerateOTPURI("MyService", secretKey)

		c.JSON(http.StatusOK, gin.H{"otp_uri": otpURI})
	})

	r.GET("/verify", func(c *gin.Context) {
		otp := c.Query("otp")
		secretKey := c.Query("secret_key")

		// Verify the OTP
		valid := VerifyOTP(otp, secretKey)

		c.JSON(http.StatusOK, gin.H{"valid": valid})
	})

	r.Run(":8080")
}
