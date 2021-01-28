package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/services/notification_service"
	"gorm.io/gorm"
)

func (r *friendRequestResolver) User(ctx context.Context, obj *models.FriendRequest) (*models.User, error) {
	return &obj.FriendRequestUser, facades.UseDB().Preload("FriendRequestUser").First(obj).Error
}

func (r *friendRequestResolver) Friend(ctx context.Context, obj *models.FriendRequest) (*models.User, error) {
	return &obj.FriendRequestFriend, facades.UseDB().Preload("FriendRequestFriend").First(obj).Error
}

func (r *mutationResolver) SendFriendRequest(ctx context.Context, userID int64) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	friend := new(models.User)
	if err := facades.UseDB().First(friend, userID).Error; err != nil {
		return nil, err
	}

	if err := notification_service.Notify(user, fmt.Sprintf("%v has requested to become your friend.", friend.DisplayName)); err != nil {
		return nil, err
	}

	return friend, facades.UseDB().Create(&models.FriendRequest{
		FriendRequestUser:   *user,
		FriendRequestFriend: *friend,
	}).Error
}

func (r *mutationResolver) AcceptFriendRequest(ctx context.Context, userID int64) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	friend := new(models.User)
	if err := facades.UseDB().First(friend, userID).Error; err != nil {
		return nil, err
	}

	return friend, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&models.Friendship{
			User:   *user,
			Friend: *friend,
		}).Error; err != nil {
			return err
		}

		return tx.Where("user_id = ?", friend.ID).
			Where("friend_id = ?", user.ID).
			Delete(new(models.FriendRequest)).
			Error
	})
}

func (r *mutationResolver) RejectFriendRequest(ctx context.Context, userID int64) (*models.User, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	friend := new(models.User)
	if err := facades.UseDB().First(friend, userID).Error; err != nil {
		return nil, err
	}

	return friend, facades.UseDB().
		Where("user_id = ?", friend.ID).
		Where("friend_id = ?", user.ID).
		Delete(new(models.FriendRequest)).
		Error
}

func (r *queryResolver) UserByFriendCode(ctx context.Context, code string) (*models.User, error) {
	user := new(models.User)
	return user, facades.UseDB().First(user, "friend_code = ?", code).Error
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

func (r *userResolver) OutgoingFriendRequests(ctx context.Context, obj *models.User) ([]*models.User, error) {
	requests := make([]*models.FriendRequest, 0)
	if err := facades.UseDB().Preload("FriendRequestUser").Where("user_id = ?", obj.ID).Find(&requests).Error; err != nil {
		return nil, err
	}

	friends := make([]*models.User, 0)
	for _, r := range requests {
		friends = append(friends, &r.FriendRequestFriend)
	}

	return friends, nil
}

func (r *userResolver) IngoingFriendRequests(ctx context.Context, obj *models.User) ([]*models.User, error) {
	requests := make([]*models.FriendRequest, 0)
	if err := facades.UseDB().Preload("FriendRequestUser").Where("friend_id = ?", obj.ID).Find(&requests).Error; err != nil {
		return nil, err
	}

	senders := make([]*models.User, 0)
	for _, r := range requests {
		senders = append(senders, &r.FriendRequestUser)
	}

	return senders, nil
}

// FriendRequest returns generated.FriendRequestResolver implementation.
func (r *Resolver) FriendRequest() generated.FriendRequestResolver { return &friendRequestResolver{r} }

type friendRequestResolver struct{ *Resolver }
