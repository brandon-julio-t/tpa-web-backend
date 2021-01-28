package notification_service

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func Notify(user *models.User, message string) error {
	return facades.UseDB().Create(&models.Notification{
		UserID:  user.ID,
		Content: message,
	}).Error
}
