package database_seeds

import (
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
)

func SeedCommunityDiscussions() error {
	for i := 0; i < 5; i++ {
		game := new(models.Game)
		if err := facades.UseDB().Order("random()").First(game).Error; err != nil {
			return err
		}

		for j := 0; j < 10; j++ {
			user := new(models.User)
			if err := facades.UseDB().Order("random()").First(user).Error; err != nil {
				return err
			}

			now := time.Now()
			aWeekAgo := now.AddDate(0, 0, -7)

			if err := facades.UseDB().
				Create(&models.CommunityDiscussion{
					Body:      faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
					CreatedAt: faker.Time().Between(aWeekAgo, now),
					Game_:     *game,
					Title:     faker.Lorem().Sentence(faker.Number().NumberInt(1)),
					User_:     *user,
				}).
				Error; err != nil {
				return err
			}
		}
	}

	return nil
}
