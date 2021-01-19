package bootstrap

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"log"
)

func init() {
	if err := facades.UseDB().
		Exec("drop table if exists game_tag_mappings").
		Exec("drop table if exists friends").
		Error; err != nil {
		log.Fatal(err)
	}

	if err := facades.UseDB().Migrator().DropTable(
		&models.Country{},
		&models.User{},
		&models.RegisterVerificationToken{},
		&models.Report{},
		&models.UnsuspendRequest{},
		&models.AssetFile{},
		&models.Game{},
		&models.GameSlideshow{},
		&models.GameTag{},
		&models.Promo{},
		&models.TopUpCode{},
		&models.ProfileComment{},
	); err != nil {
		log.Fatal(err)
	}

	if err := facades.UseDB().AutoMigrate(
		&models.Country{},
		&models.User{},
		&models.RegisterVerificationToken{},
		&models.Report{},
		&models.UnsuspendRequest{},
		&models.AssetFile{},
		&models.Game{},
		&models.GameSlideshow{},
		&models.GameTag{},
		&models.Promo{},
		&models.TopUpCode{},
		&models.ProfileComment{},
	); err != nil {
		log.Fatal(err)
	}
}
