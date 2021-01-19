package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
)

func (r *mutationResolver) Befriend(ctx context.Context, userID int64) (*models.User, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return nil, errors.New("not authenticated")
	}

	var friend models.User
	if err := facades.UseDB().First(&friend, userID).Error; err != nil {
		return nil, err
	}

	return &friend, facades.UseDB().Model(&user).Association("Friends").Append(&friend)
}

func (r *mutationResolver) Unfriend(ctx context.Context, userID int64) (*models.User, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return nil, errors.New("not authenticated")
	}

	var friend models.User
	if err := facades.UseDB().First(&friend, userID).Error; err != nil {
		return nil, err
	}

	return &friend, facades.UseDB().Model(&user).Association("Friends").Delete(&friend)
}

func (r *queryResolver) Friends(ctx context.Context) ([]*models.User, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return nil, errors.New("not authenticated")
	}

	var friends []*models.User
	return friends, facades.UseDB().Model(&user).Association("Friends").Find(&friends)
}
