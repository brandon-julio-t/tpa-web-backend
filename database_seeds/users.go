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
			File:        defaultProfilePicture,
			ContentType: "image/png",
		},
		WalletBalance: 100000,
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
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "BR",
			Email:       "brandon.julio.t@icloud.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "CC",
			Email:       "cc@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "ST",
			Email:       "st@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "VN",
			Email:       "vn@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "TC",
			Email:       "tc@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "LL",
			Email:       "ll@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "GA",
			Email:       "ga@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "JP",
			Email:       "jp@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		},
		{
			AccountName: "AE",
			Email:       "ae@slc.com",
			Password:    string(userHash),
			CountryID:   69,
			ProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
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
