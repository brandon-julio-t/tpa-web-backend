package bootstrap

import (
	"log"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func init() {
	if err := facades.UseDB().
		Exec("drop table if exists game_tag_mappings cascade").
		Exec("drop table if exists friends cascade").
		Exec("drop table if exists wishlist cascade").
		Exec("drop table if exists cart cascade").
		Exec("drop table if exists friend_requests cascade").
		Error; err != nil {
		log.Fatal(err)
	}

	entities := []interface{}{
		new(models.AssetFile),
		new(models.Country),
		new(models.Promo),
		new(models.RegisterVerificationToken),
		new(models.Report),
		new(models.TopUpCode),
		new(models.UnsuspendRequest),

		new(models.Game),
		new(models.GameGenre),
		new(models.GameSlideshow),
		new(models.GameTag),
		new(models.GameReview),
		new(models.GameReviewVote),
		new(models.GameReviewComment),

		new(models.User),
		new(models.Friendship),
		new(models.FriendRequest),
		new(models.PrivateMessage),
		new(models.ProfileComment),
		new(models.Notification),

		new(models.MarketItem),
		new(models.MarketItemOffer),
		new(models.MarketItemTransaction),

		new(models.Inventory),

		new(models.GamePurchaseTransactionHeader),
		new(models.GamePurchaseTransactionDetail),
		new(models.GameGiftTransactionHeader),
		new(models.GameGiftTransactionDetail),

		new(models.CommunityImageAndVideo),
		new(models.CommunityImageAndVideoComment),
		new(models.CommunityImageAndVideoRating),
		new(models.CommunityDiscussion),
		new(models.CommunityDiscussionComment),

		new(models.PointItem),
		new(models.PointItemPurchase),
	}

	if err := facades.UseDB().Migrator().DropTable(entities...); err != nil {
		log.Fatal(err)
	}

	if err := facades.UseDB().AutoMigrate(entities...); err != nil {
		log.Fatal(err)
	}
}
