package facades

import (
	"github.com/google/uuid"
	"math/rand"
)

func UseOTP() string {
	t1, t2 := uuid.Must(uuid.NewRandom()).Time().UnixTime()

	rand.Seed(t1 + t2)

	alphanumeric := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	otp := make([]rune, 5)

	for i, _ := range otp {
		otp[i] = rune(alphanumeric[rand.Intn(len(alphanumeric))])
	}

	return string(otp)
}
