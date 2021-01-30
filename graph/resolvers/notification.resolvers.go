package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"

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

func (r *userResolver) Notifications(ctx context.Context, obj *models.User) ([]*models.Notification, error) {
	notifications := make([]*models.Notification, 0)
	return notifications, facades.UseDB().Where("user_id = ?", obj.ID).Find(&notifications).Error
}

func (r *userResolver) ReceivedProfileCommentsCount(ctx context.Context, obj *models.User) (int64, error) {
	key := fmt.Sprintf("notification:profile_comments:%v", obj.ID)

	cached, err := facades.UseCached(ctx, key, func() (string, error) {
		count := new(int64)

		if err := facades.UseDB().
			Model(new(models.Notification)).
			Where("user_id = ?", obj.ID).
			Where("content like '%profile%'").
			Count(count).
			Error; err != nil {
			return "-1", err
		}

		return strconv.Itoa(int(*count)), nil
	})

	if err != nil {
		return -1, err
	}

	return strconv.ParseInt(cached, 10, 64)
}

func (r *userResolver) ReceivedInvitesCount(ctx context.Context, obj *models.User) (int64, error) {
	key := fmt.Sprintf("notification:invites:%v", obj.ID)

	cached, err := facades.UseCached(ctx, key, func() (string, error) {
		count := new(int64)

		if err := facades.UseDB().
			Model(new(models.Notification)).
			Where("user_id = ?", obj.ID).
			Where("content like '%friend%'").
			Count(count).
			Error; err != nil {
			return "-1", err
		}

		return strconv.Itoa(int(*count)), nil
	})

	if err != nil {
		return -1, err
	}

	return strconv.ParseInt(cached, 10, 64)
}

func (r *userResolver) ReceivedGiftsCount(ctx context.Context, obj *models.User) (int64, error) {
	key := fmt.Sprintf("notification:gifts:%v", obj.ID)

	cached, err := facades.UseCached(ctx, key, func() (string, error) {
		count := new(int64)

		if err := facades.UseDB().
			Model(new(models.Notification)).
			Where("user_id = ?", obj.ID).
			Where("content like '%gift%'").
			Count(count).
			Error; err != nil {
			return "-1", err
		}

		return strconv.Itoa(int(*count)), nil
	})

	if err != nil {
		return -1, err
	}

	return strconv.ParseInt(cached, 10, 64)
}

func (r *userResolver) ReceivedMessagesCount(ctx context.Context, obj *models.User) (int64, error) {
	key := fmt.Sprintf("notification:messages:%v", obj.ID)

	cached, err := facades.UseCached(ctx, key, func() (string, error) {
		count := new(int64)

		if err := facades.UseDB().
			Model(new(models.Notification)).
			Where("user_id = ?", obj.ID).
			Where("content like '%message%'").
			Count(count).
			Error; err != nil {
			return "-1", err
		}

		return strconv.Itoa(int(*count)), nil
	})

	if err != nil {
		return -1, err
	}

	return strconv.ParseInt(cached, 10, 64)
}
