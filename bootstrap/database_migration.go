package bootstrap

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"log"
)

func init() {
	if err := facades.UseDB().
		Exec("drop table if exists game_tag_mappings cascade").
		Exec("drop table if exists friends cascade").
		Exec("drop table if exists wishlist cascade").
		Exec("drop table if exists cart cascade").
		Error; err != nil {
		log.Fatal(err)
	}

	if err := facades.UseDB().Migrator().DropTable(
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
		new(models.ProfileComment),

		new(models.User),
		new(models.PrivateMessage),
		new(models.Friendship),

		new(models.GamePurchaseTransactionHeader),
		new(models.GamePurchaseTransactionDetail),
		new(models.GameGiftTransactionHeader),
		new(models.GameGiftTransactionDetail),
	); err != nil {
		log.Fatal(err)
	}

	if err := facades.UseDB().AutoMigrate(
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

		new(models.User),
		new(models.Friendship),
		new(models.PrivateMessage),
		new(models.ProfileComment),

		new(models.GamePurchaseTransactionHeader),
		new(models.GamePurchaseTransactionDetail),
		new(models.GameGiftTransactionHeader),
		new(models.GameGiftTransactionDetail),

	); err != nil {
		log.Fatal(err)
	}
}
