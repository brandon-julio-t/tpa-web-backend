package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
)

func SeedGameReviewVotes() error {
	users := make([]*models.User, 0)
	if err := facades.UseDB().Find(&users).Error; err != nil {
		return err
	}

	reviews := make([]*models.GameReview, 0)
	if err := facades.UseDB().Find(&reviews).Error; err != nil {
		return err
	}

	for _, user := range users {
		for _, review := range reviews {
			facades.UseDB().Create(&models.GameReviewVote{
				GameReviewVoteGameReview: *review,
				GameReviewVoteUser:       *user,
				IsUpVote:                 faker.Number().NumberInt(1)%2 == 0,
			})
		}
	}

	return nil
}
