package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
	"time"
)

func SeedCommunityImageAndVideoComments() error {
	imagesAndVideos := make([]*models.CommunityImageAndVideo, 0)
	if err := facades.UseDB().Find(&imagesAndVideos).Error; err != nil {
		return err
	}

	for _, imgAndVid := range imagesAndVideos {
		for i := 0; i < 20; i++ {
			now := time.Now()
			before := now.AddDate(0, 0, faker.Number().NumberInt(1)*-1)

			user := new(models.User)
			if err := facades.UseDB().Order("random()").First(user).Error; err != nil {
				return err
			}

			if err := facades.UseDB().Create(&models.CommunityImageAndVideoComment{
				Body:                    faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
				CreatedAt:               faker.Time().Between(before, now),
				User_:                   *user,
				CommunityImageAndVideo_: *imgAndVid,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
