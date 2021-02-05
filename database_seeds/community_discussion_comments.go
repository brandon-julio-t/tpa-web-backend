package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
	"time"
)

func SeedCommunityDiscussionComments() error {
	discussions := make([]*models.CommunityDiscussion, 0)
	if err := facades.UseDB().Find(&discussions).Error; err != nil {
		return err
	}

	for _, discussion := range discussions {
		for i := 0; i < 20; i++ {
			user := new(models.User)
			if err := facades.UseDB().Order("random()").First(user).Error; err != nil {
				return err
			}

			now := time.Now()
			before := now.AddDate(0, 0, faker.Number().NumberInt(1)*-1)

			if err := facades.UseDB().Create(&models.CommunityDiscussionComment{
				Body:                 faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
				CreatedAt:            faker.Time().Between(before, now),
				CommunityDiscussion_: *discussion,
				User_:                *user,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
