package facades

import (
	"math/rand"
	"time"
)

func UseOTP() string {
	rand.Seed(time.Now().Unix())

	alphanumeric := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	otp := make([]rune, 5)

	for i, _ := range otp {
		otp[i] = rune(alphanumeric[rand.Intn(len(alphanumeric))])
	}

	return string(otp)
}
