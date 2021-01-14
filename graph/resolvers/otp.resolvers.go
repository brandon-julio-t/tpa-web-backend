package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *mutationResolver) SendOtp(ctx context.Context, email string) (*bool, error) {
	sender := os.Getenv("MAILTRAP_EMAIL")
	receivers := []string{email}

	otp := facades.UseOTP()
	if err := facades.UseDB().Create(&models.RegisterVerificationToken{Token: otp}).Error; err != nil {
		return nil, err
	}

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Register</title>
</head>
<body>
	Your STAEM registration OTP: %v
</body>
</html>
`, otp)

	message :=
		fmt.Sprintf(
			"From: %v\r\n"+
				"To: %v\r\n"+
				"Subject: Testing\r\n"+
				"MIME-version: 1.0;\r\n"+
				"Content-Type: text/html; charset=\"UTF-8\";\r\n"+
				"\r\n"+
				"%v\r\n",
			sender,
			strings.Join(receivers, ","),
			html,
		)

	if err := smtp.SendMail(
		os.Getenv("MAILTRAP_ADDRESS"),
		smtp.PlainAuth("", os.Getenv("MAILTRAP_USERNAME"), os.Getenv("MAILTRAP_PASSWORD"), os.Getenv("MAILTRAP_HOST")),
		sender,
		receivers,
		[]byte(message)); err != nil {
		return nil, err
	}

	result := true
	return &result, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, otp string) (*bool, error) {
	var registerVerificationToken models.RegisterVerificationToken
	if err := facades.UseDB().First(&registerVerificationToken, "token = ?", otp).Error; err != nil {
		return nil, err
	}

	result := true

	if err := facades.UseDB().Delete(&registerVerificationToken).Error; err != nil {
		return nil, err
	}

	return &result, nil
}
