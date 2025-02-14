package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/services/notification_service"
)

func (r *mutationResolver) CreateProfileComment(ctx context.Context, profileID int64, comment string) (*models.ProfileComment, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	profile := new(models.User)
	if err := facades.UseDB().First(profile, profileID).Error; err != nil {
		return nil, err
	}

	profileComment := models.ProfileComment{
		UserID:    user.ID,
		ProfileID: profileID,
		Comment:   comment,
	}

	if err := facades.UseDB().Create(&profileComment).Error; err != nil {
		return nil, err
	}

	if err := notification_service.Notify(profile, fmt.Sprintf("%v commented on your profile", user.DisplayName)); err != nil {
		return nil, err
	}

	return &profileComment, facades.UseDB().
		Preload("User").
		Preload("User.UserProfilePicture").
		First(&profileComment).
		Error
}

func (r *mutationResolver) DeleteProfileComment(ctx context.Context, id int64) (*models.ProfileComment, error) {
	profileComment := models.ProfileComment{ID: id}
	return &profileComment, facades.UseDB().Delete(&profileComment, id).Error
}

func (r *queryResolver) ProfileComments(ctx context.Context, profileID int64) ([]*models.ProfileComment, error) {
	var profileComments []*models.ProfileComment
	return profileComments, facades.UseDB().
		Preload("User").
		Preload("User.UserProfilePicture").
		Where("profile_id = ?", profileID).
		Order("created_at desc").
		Find(&profileComments).
		Error
}
