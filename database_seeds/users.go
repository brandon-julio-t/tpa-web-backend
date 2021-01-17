package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func SeedUsers() error {
	adminHash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Omit("ProfilePicture").Create(&models.User{
		AccountName: "Admin",
		Email:       "admin@admin.com",
		Password:    string(adminHash),
		CountryID:   69,
		ProfilePictureID: 10,
	})

	userHash, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	for _, user := range []*models.User{
		{AccountName: "User", Email: "user@user.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 1},
		{AccountName: "BR", Email: "br@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 2},
		{AccountName: "CC", Email: "cc@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 3},
		{AccountName: "ST", Email: "st@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 4},
		{AccountName: "VN", Email: "vn@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 5},
		{AccountName: "TC", Email: "tc@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 6},
		{AccountName: "LL", Email: "ll@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 7},
		{AccountName: "GA", Email: "ga@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 8},
		{AccountName: "JP", Email: "jp@slc.com", Password: string(userHash), CountryID: 69, ProfilePictureID: 9},
	} {
		facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Omit("ProfilePicture").Create(user)
	}

	return nil
}
