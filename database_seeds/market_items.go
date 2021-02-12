package database_seeds

import (
	"io/ioutil"
	"path/filepath"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"syreclabs.com/go/faker"
)

func SeedMarketItems() error {
	imagePath := filepath.Join("assets", "Background Zoom SLC.png")
	img, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return err
	}

	games := make([]*models.Game, 0)
	if err := facades.UseDB().Find(&games).Error; err != nil {
		return err
	}

	for _, game := range games {
		for i := 0; i < 20; i++ {
			if err := facades.UseDB().Create(&models.MarketItem{
				Description: faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
				Game_:       *game,
				ImageRef: models.AssetFile{
					File:        img,
					ContentType: "image/png",
				},
				Name: faker.App().Name(),
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
