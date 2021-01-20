package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"path/filepath"
)

func SeedUsers() error {
	path := filepath.Join("assets", "default_profile_picture.png")
	defaultProfilePicture, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	adminHash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(&models.User{
		AccountName: "Admin",
		Email:       "admin@admin.com",
		Password:    string(adminHash),
		CountryID:   69,
		ProfilePicture: models.AssetFile{
			ID:          1,
			File:        defaultProfilePicture,
			ContentType: "image/png",
		},
	})

	userHash, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users := []*models.User{
		{
			AccountName: "User",
			Email:       "user@user.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          2,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "BR",
			Email:       "br@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          3,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "CC",
			Email:       "cc@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          4,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "ST",
			Email:       "st@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          5,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "VN",
			Email:       "vn@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          6,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "TC",
			Email:       "tc@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          7,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "LL",
			Email:       "ll@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          8,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "GA",
			Email:       "ga@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          9,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
		{
			AccountName: "JP",
			Email:       "jp@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				ID:          10,
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
		},
	}

	for _, user := range users {
		facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(user)
	}

	for _, user := range users {
		for _, friend := range users {
			if user.ID == friend.ID {
				continue
			}

			if err := facades.UseDB().Create(&models.Friendship{
				UserID:   user.ID,
				User:     *user,
				FriendID: friend.ID,
				Friend:   *friend,
			}).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
