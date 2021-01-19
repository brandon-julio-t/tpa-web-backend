package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *mutationResolver) CreatePromo(ctx context.Context, discount float64, endAt time.Time) (*models.Promo, error) {
	promo := models.Promo{
		Discount: discount,
		EndAt:    endAt,
	}

	return &promo, facades.UseDB().Create(&promo).Error
}

func (r *mutationResolver) UpdatePromo(ctx context.Context, id int64, discount float64, endAt time.Time) (*models.Promo, error) {
	var promo models.Promo
	if err := facades.UseDB().First(&promo, id).Error; err != nil {
		return nil, err
	}
	promo.Discount = discount
	promo.EndAt = endAt
	return &promo, facades.UseDB().Save(&promo).Error
}

func (r *mutationResolver) DeletePromo(ctx context.Context, id int64) (*models.Promo, error) {
	var promo models.Promo
	if err := facades.UseDB().First(&promo, id).Error; err != nil {
		return nil, err
	}
	return &promo, facades.UseDB().Delete(&promo).Error
}

func (r *queryResolver) Promos(ctx context.Context, page int) (*models.PromoPagination, error) {
	var promos []*models.Promo
	var count int64
	pageSize := 5

	if err := facades.UseDB().
		Model(&models.Promo{}).
		Count(&count).
		Scopes(facades.UsePagination(page, pageSize)).
		Find(&promos).Error; err != nil {
		return nil, err
	}

	return &models.PromoPagination{
		Data:       promos,
		TotalPages: int(math.Ceil(float64(count) / float64(pageSize))),
	}, nil
}

func (r *queryResolver) Promo(ctx context.Context, id int64) (*models.Promo, error) {
	var promo models.Promo
	return &promo, facades.UseDB().First(&promo, id).Error
}
