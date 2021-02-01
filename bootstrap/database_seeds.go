package bootstrap

import (
	"github.com/brandon-julio-t/tpa-web-backend/database_seeds"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"gorm.io/gorm"
	"log"
)

func init() {
	if err := facades.UseDB().Transaction(func(tx *gorm.DB) error {
		for _, seed := range []func() error{
			database_seeds.SeedCountries,
			database_seeds.SeedUsers,
			database_seeds.SeedGameGenres,
			database_seeds.SeedGameTags,
			database_seeds.SeedPromos,
			database_seeds.SeedTopUpCodes,
			database_seeds.SeedGames,
			database_seeds.SeedGameReviewVotes,
			database_seeds.SeedGamePurchaseTransactions,
			database_seeds.SeedGameGiftTransactions,
			database_seeds.SeedCommunityImagesAndVideos,
			database_seeds.SeedCommunityImageAndVideoComments,
			database_seeds.SeedCommunityImageAndVideoDislikesLikes,
		} {
			if err := seed(); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
