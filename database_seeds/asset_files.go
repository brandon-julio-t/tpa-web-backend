package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"path/filepath"
)

func SeedAssetFiles() error {
	path := filepath.Join("assets", "default_profile_picture.png")
	defaultProfilePicture, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	for i := 1; i <= 10; i++ {
		if err := facades.UseDB().
			Clauses(clause.OnConflict{DoNothing: true}).
			Create(&models.AssetFile{ID: int64(i), File: defaultProfilePicture, ContentType: "image/png"}).
			Error; err != nil {
			return err
		}
	}

	return nil
}
