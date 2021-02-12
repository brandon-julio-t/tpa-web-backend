package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/repositories"
	"gorm.io/gorm"
)

func (r *mutationResolver) PurchasePointItem(ctx context.Context, id int64) (*models.PointItem, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	item := new(models.PointItem)
	if err := facades.UseDB().First(item, id).Error; err != nil {
		return nil, err
	}

	return item, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		user.Points -= int64(item.Price)
		if err := tx.Save(user).Error; err != nil {
			return err
		}

		return tx.Create(&models.PointItemPurchase{
			User_:      *user,
			PointItem_: *item,
		}).Error
	})
}

func (r *mutationResolver) EditAvatarBorder(ctx context.Context, id int64) (*models.PointItem, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	item, err := new(repositories.PointItemRepository).GetByIDAndCategory(id, "avatar_border")
	if err != nil {
		return nil, err
	}

	user.AvatarBorderID = item.ID
	return item, facades.UseDB().Save(user).Error
}

func (r *mutationResolver) EditProfileBackground(ctx context.Context, id int64) (*models.PointItem, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	item, err := new(repositories.PointItemRepository).GetByIDAndCategory(id, "profile_background")
	if err != nil {
		return nil, err
	}

	user.ProfileBackgroundID = item.ID
	return item, facades.UseDB().Save(user).Error
}

func (r *mutationResolver) EditMiniProfileBackground(ctx context.Context, id int64) (*models.PointItem, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	item, err := new(repositories.PointItemRepository).GetByIDAndCategory(id, "mini_profile_background")
	if err != nil {
		return nil, err
	}

	user.MiniProfileBackgroundID = item.ID
	return item, facades.UseDB().Save(user).Error
}

func (r *queryResolver) PointItemProfileBackgrounds(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetProfileBackgrounds()
}

func (r *queryResolver) PointItemAvatarBorders(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetAvatarBorders()
}

func (r *queryResolver) PointItemAnimatedAvatars(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetAnimatedAvatars()
}

func (r *queryResolver) PointItemChatStickers(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetChatStickers()
}

func (r *queryResolver) PointItemMiniProfileBackgrounds(ctx context.Context) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetMiniProfileBackgrounds()
}

func (r *userResolver) AvatarBorder(ctx context.Context, obj *models.User) (*models.PointItem, error) {
	if obj.AvatarBorderID == 0 {
		return nil, nil
	}
	return new(repositories.PointItemRepository).GetByIDAndCategory(obj.AvatarBorderID, "avatar_border")
}

func (r *userResolver) MiniProfileBackground(ctx context.Context, obj *models.User) (*models.PointItem, error) {
	if obj.MiniProfileBackgroundID == 0 {
		return nil, nil
	}
	return new(repositories.PointItemRepository).GetByIDAndCategory(obj.MiniProfileBackgroundID, "mini_profile_background")
}

func (r *userResolver) OwnedAvatarBorders(ctx context.Context, obj *models.User) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetAllByUserAndCategory(obj, "avatar_border")
}

func (r *userResolver) OwnedProfileBackgrounds(ctx context.Context, obj *models.User) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetAllByUserAndCategory(obj, "profile_background")
}

func (r *userResolver) OwnedMiniProfileBackgrounds(ctx context.Context, obj *models.User) ([]*models.PointItem, error) {
	return new(repositories.PointItemRepository).GetAllByUserAndCategory(obj, "mini_profile_background")
}

func (r *userResolver) ProfileBackground(ctx context.Context, obj *models.User) (*models.PointItem, error) {
	if obj.ProfileBackgroundID == 0 {
		return nil, nil
	}
	return new(repositories.PointItemRepository).GetByIDAndCategory(obj.ProfileBackgroundID, "profile_background")
}
