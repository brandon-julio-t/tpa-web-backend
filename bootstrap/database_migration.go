package bootstrap

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"log"
)

func init() {
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
	); err != nil {
		log.Fatal(err)
	}
}
