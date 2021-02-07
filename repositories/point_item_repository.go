package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

type PointItemRepository struct{}

func (PointItemRepository) GetByIDAndCategory(id int64, category string) (*models.PointItem, error) {
	item := new(models.PointItem)
	return item, facades.UseDB().
		Where("category = ?", category).
		First(item, id).
		Error
}

func (PointItemRepository) GetAllByUserAndCategory(user *models.User, category string) ([]*models.PointItem, error) {
	purchasedItems := make([]*models.PointItemPurchase, 0)
	if err := facades.UseDB().
		Joins("join point_items pi on pi.id = point_item_purchases.point_item_id").
		Where("user_id = ?", user.ID).
		Where("category = ?", category).
		Find(&purchasedItems).
		Error; err != nil {
		return nil, err
	}

	items := make([]*models.PointItem, 0)
	for _, purchased := range purchasedItems {
		if err := facades.UseDB().Preload("Image_").First(&purchased.PointItem_).Error; err != nil {
			return nil, err
		}

		items = append(items, &purchased.PointItem_)
	}

	return items, nil
}

func (PointItemRepository) GetProfileBackgrounds() ([]*models.PointItem, error) {
	items := make([]*models.PointItem, 0)
	return items, facades.UseDB().Find(&items, "category = ?", "profile_background").Error
}

func (PointItemRepository) GetAvatarBorders() ([]*models.PointItem, error) {
	items := make([]*models.PointItem, 0)
	return items, facades.UseDB().Find(&items, "category = ?", "avatar_border").Error
}

func (PointItemRepository) GetAnimatedAvatars() ([]*models.PointItem, error) {
	items := make([]*models.PointItem, 0)
	return items, facades.UseDB().Find(&items, "category = ?", "animated_avatar").Error
}

func (PointItemRepository) GetChatStickers() ([]*models.PointItem, error) {
	items := make([]*models.PointItem, 0)
	return items, facades.UseDB().Find(&items, "category = ?", "sticker").Error
}

func (PointItemRepository) GetMiniProfileBackgrounds() ([]*models.PointItem, error) {
	items := make([]*models.PointItem, 0)
	return items, facades.UseDB().Find(&items, "category = ?", "mini_profile_background").Error
}
