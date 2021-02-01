package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
	"time"
)

func SeedCommunityImageAndVideoDislikesLikes() error {
	imagesAndVideos := make([]*models.CommunityImageAndVideo, 0)
	if err := facades.UseDB().Find(&imagesAndVideos).Error; err != nil {
		return err
	}

	users := make([]*models.User, 0)
	if err := facades.UseDB().Find(&users).Error; err != nil {
		return err
	}

	now := time.Now()
	before := now.AddDate(0, 0, faker.Number().NumberInt(1)*-1)

	for _, user := range users {
		for _, imgAndVid := range imagesAndVideos {
			if err := facades.UseDB().Create(&models.CommunityImageAndVideoRating{
				CreatedAt:               faker.Time().Between(before, now),
				IsLike:                  faker.Number().NumberInt(1)%2 == 0,
				CommunityImageAndVideo_: *imgAndVid,
				User_:                   *user,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
