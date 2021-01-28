package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *mutationResolver) DeleteNotification(ctx context.Context, id int64) (*models.Notification, error) {
	notification := new(models.Notification)
	if err := facades.UseDB().First(notification, id).Error; err != nil {
		return nil, err
	}
	return notification, facades.UseDB().Delete(notification).Error
}

func (r *queryResolver) NotificationByID(ctx context.Context, id int64) (*models.Notification, error) {
	notification := new(models.Notification)
	return notification, facades.UseDB().First(notification, id).Error
}

func (r *userResolver) Notification(ctx context.Context, obj *models.User) ([]*models.Notification, error) {
	notifications := make([]*models.Notification, 0)
	return notifications, facades.UseDB().Where("user_id = ?", obj.ID).Find(&notifications).Error
}
