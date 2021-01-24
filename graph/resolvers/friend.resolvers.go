package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
)

func (r *mutationResolver) Befriend(ctx context.Context, userID int64) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	var friend models.User
	if err := facades.UseDB().First(&friend, userID).Error; err != nil {
		return nil, err
	}

	return &friend, facades.UseDB().Create([]*models.Friendship{
		{
			UserID:   user.ID,
			User:     *user,
			FriendID: friend.ID,
			Friend:   friend,
		},
		{
			UserID:   friend.ID,
			User:     friend,
			FriendID: user.ID,
			Friend:   *user,
		},
	}).Error
}

func (r *mutationResolver) Unfriend(ctx context.Context, userID int64) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	var friend models.User
	if err := facades.UseDB().First(&friend, userID).Error; err != nil {
		return nil, err
	}

	return &friend, facades.UseDB().
		Where("user_id = ? and friend_id = ?", friend.ID, user.ID).
		Where("user_id = ? and friend_id = ?", user.ID, friend.ID).
		Delete(&models.Friendship{}).
		Error
}

func (r *userResolver) Friends(ctx context.Context, obj *models.User) ([]*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	var friendships []*models.Friendship
	if err := facades.UseDB().
		Preload("User").
		Preload("User.UserProfilePicture").
		Preload("Friend").
		Preload("Friend.UserProfilePicture").
		Find(&friendships, "user_id = ?", user.ID).
		Error; err != nil {
		return nil, err
	}

	var friends []*models.User

	for _, friendship := range friendships {
		friends = append(friends, &friendship.Friend)
	}

	return friends, nil
}
