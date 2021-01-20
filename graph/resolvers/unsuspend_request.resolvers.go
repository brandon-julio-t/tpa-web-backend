package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
)

func (r *mutationResolver) UnsuspendRequest(ctx context.Context, accountName string) (string, error) {
	repo := new(repositories.UserRepository)

	user, err := repo.GetByAccountName(accountName)
	if err != nil {
		return "", err
	}

	if err := facades.UseDB().Create(&models.UnsuspendRequest{
		UserID: user.ID,
		User:   *user,
	}).Error; err != nil {
		return "", err
	}

	return user.AccountName, nil
}

func (r *mutationResolver) ApproveUnsuspendRequests(ctx context.Context, id int64) (*models.User, error) {
	repo := new(repositories.UserRepository)

	user, err := repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	user.SuspendedAt = time.Time{}

	if _, err := repo.Update(user); err != nil {
		return nil, err
	}

	if err := facades.UseDB().Delete(&models.UnsuspendRequest{}, "user_id = ?", id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *mutationResolver) DenyUnsuspendRequests(ctx context.Context, id int64) (*models.User, error) {
	if err := facades.UseDB().Delete(&models.UnsuspendRequest{}, "user_id = ?", id).Error; err != nil {
		return nil, err
	}

	return new(repositories.UserRepository).GetByID(id)
}

func (r *queryResolver) GetAllUnsuspendRequests(ctx context.Context) ([]*models.User, error) {
	var unsuspendRequests []*models.UnsuspendRequest
	if err := facades.UseDB().Preload("Sender").Find(&unsuspendRequests).Error; err != nil {
		return nil, err
	}

	var users []*models.User
	for _, x := range unsuspendRequests {
		users = append(users, &x.User)
	}

	return users, nil
}
