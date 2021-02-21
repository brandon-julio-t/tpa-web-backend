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
		},
		{
			AccountName: "BR",
			Email:       "brandon.julio.t@icloud.com",
		},
		{
			AccountName: "CC",
			Email:       "cc@slc.com",
		},
		{
			AccountName: "ST",
			Email:       "st@slc.com",
		},
		{
			AccountName: "VN",
			Email:       "vn@slc.com",
		},
		{
			AccountName: "TC",
			Email:       "tc@slc.com",
		},
		{
			AccountName: "LL",
			Email:       "ll@slc.com",
		},
		{
			AccountName: "GA",
			Email:       "ga@slc.com",
		},
		{
			AccountName: "JP",
			Email:       "jp@slc.com",
		},
		{
			AccountName: "AE",
			Email:       "ae@slc.com",
		},
	}

	for _, user := range users {
		country := new(models.Country)
		if err := facades.UseDB().Order("random()").First(country).Error; err != nil {
			return err
		}

		facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(&models.User{
			AccountName: user.AccountName,
			Country:     *country,
			Email:       user.Email,
			Exp:         faker.Number().NumberInt64(faker.Number().NumberInt(1)),
			Password:    string(userHash),
			Points:      100000,
			Status:      faker.RandomChoice([]string{"online", "offline", "playing"}),
			Summary:     faker.Lorem().Paragraph(faker.Number().NumberInt(1)),
			UserProfilePicture: models.AssetFile{
				File:        defaultProfilePicture,
				ContentType: "image/png",
			},
			WalletBalance: 100000,
		})
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
