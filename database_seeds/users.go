package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"path/filepath"
	"syreclabs.com/go/faker"
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
		Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
		Password:    string(adminHash),
		CountryID:   69,
		UserProfilePicture: models.AssetFile{
			File:        defaultProfilePicture,
			ContentType: "image/png",
		},
		Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
		Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
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
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "BR",
			Email:       "brandon.julio.t@icloud.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "CC",
			Email:       "cc@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "ST",
			Email:       "st@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "VN",
			Email:       "vn@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "TC",
			Email:       "tc@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "LL",
			Email:       "ll@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "GA",
			Email:       "ga@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "JP",
			Email:       "jp@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
		{
			AccountName: "AE",
			Email:       "ae@slc.com",
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			CountryID:   69,
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			Status:        faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:       faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			WalletBalance: 100000,
		},
	}

	for _, user := range users {
		facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(user)
	}

	for _, user := range users {
		count := 0

		for _, friend := range users {
			if count > 4 || user.ID == friend.ID {
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

			count++
		}
	}

	return nil
}
