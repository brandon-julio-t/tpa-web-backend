package bootstrap

import (
	"log"

	"github.com/brandon-julio-t/tpa-web-backend/database_seeds"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"gorm.io/gorm"
)

func init() {
	go func() {
		if err := facades.UseDB().Transaction(func(tx *gorm.DB) error {
			for _, seed := range []func() error{
				database_seeds.SeedCountries,
				database_seeds.SeedUsers,
				database_seeds.SeedFriends,
				database_seeds.SeedPromos,
				database_seeds.SeedTopUpCodes,

				database_seeds.SeedGameGenres,
				database_seeds.SeedGameTags,
				database_seeds.SeedGames,
				database_seeds.SeedGameReviewVotes,
				database_seeds.SeedGameReviewComments,
				database_seeds.SeedGamePurchaseTransactions,
				database_seeds.SeedGameGiftTransactions,

				database_seeds.SeedCommunityImagesAndVideos,
				database_seeds.SeedCommunityImageAndVideoComments,
				database_seeds.SeedCommunityImageAndVideoDislikesLikes,
				database_seeds.SeedCommunityDiscussions,
				database_seeds.SeedCommunityDiscussionComments,

				database_seeds.SeedPointItems,
			} {
				if err := seed(); err != nil {
					return err
				}
			}

			return nil
		}); err != nil {
			log.Fatal(err)
		}

		log.Print("Database seeding finished.")
	}()
}
