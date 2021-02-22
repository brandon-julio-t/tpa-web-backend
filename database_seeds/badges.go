package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
)

func SeedBadges() error {
	games := make([]*models.Game, 0)
	if err := facades.UseDB().Find(&games).Error; err != nil {
		return err
	}

	for _, game := range games {
		badge := &models.Badge{
			Exp:   100,
			Game:  *game,
			Level: 1,
			Name:  faker.Lorem().Word(),
		}

		if err := facades.UseDB().Create(badge).Error; err != nil {
			return err
		}

		for i := 0; i < 5; i++ {
			if err := facades.UseDB().Create(&models.BadgeCard{
				Badge:   *badge,
				Name:    faker.Lorem().Word(),
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
