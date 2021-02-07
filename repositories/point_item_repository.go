package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

type PointItemRepository struct{}

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
