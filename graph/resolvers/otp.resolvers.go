package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/mailjet/mailjet-apiv3-go"
	"os"
)

func (r *mutationResolver) SendOtp(ctx context.Context, email string) (bool, error) {
	otp := facades.UseOTP()

	if err := facades.UseDB().Create(&models.RegisterVerificationToken{Token: otp}).Error; err != nil {
		return false, err
	}

	mailjetClient := mailjet.NewMailjetClient(os.Getenv("MAILJET_PUBLIC_KEY"), os.Getenv("MAILJET_PRIVATE_KEY"))
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: "brandon.julio.t@icloud.com",
				Name:  "STAEM",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: email,
				},
			},
			Subject:  "STAEM create account OTP",
			TextPart: fmt.Sprintf("Your Staem registration OTP: %v", otp),
			CustomID: "StaemCreateAccountOTP",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	if _, err := mailjetClient.SendMailV31(&messages); err != nil {
		return false, err
	}

	return true, nil
}

func (r *mutationResolver) VerifyOtp(ctx context.Context, otp string) (bool, error) {
	var registerVerificationToken models.RegisterVerificationToken
	if err := facades.UseDB().First(&registerVerificationToken, "token = ?", otp).Error; err != nil {
		return false, err
	}

	if err := facades.UseDB().Delete(&registerVerificationToken).Error; err != nil {
		return false, err
	}

	return true, nil
}
