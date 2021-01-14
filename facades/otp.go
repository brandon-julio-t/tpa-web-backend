package facades

import (
	"math/rand"
	"time"
)

const (
	min = 33
	max = 126
)

func UseOTP() string {
	rand.Seed(time.Now().Unix())
	otp := make([]rune, 5)
	for i, _ := range otp {
		otp[i] = rune(rand.Intn(max-min+1) + min)
	}
	return string(otp)
}
