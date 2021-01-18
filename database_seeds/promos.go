package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"time"
)

func SeedPromos() error {
	promos := []*models.Promo{
		{Discount: 10, EndAt: time.Now()},
		{Discount: 15, EndAt: time.Now()},
		{Discount: 20, EndAt: time.Now()},
		{Discount: 25, EndAt: time.Now()},
		{Discount: 30, EndAt: time.Now()},
		{Discount: 35, EndAt: time.Now()},
		{Discount: 40, EndAt: time.Now()},
		{Discount: 45, EndAt: time.Now()},
		{Discount: 50, EndAt: time.Now()},
	}
	return facades.UseDB().Create(&promos).Error
}
