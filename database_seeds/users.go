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

	facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(&models.User{
		AccountName: "Admin",
		Email:       "admin@admin.com",
		Password:    string(adminHash),
		CountryID:   69,
	})

	userHash, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	for _, user := range []*models.User{
		{AccountName: "User", Email: "user@user.com", Password: string(userHash), CountryID: 420},
		{AccountName: "BR", Email: "br@slc.com", Password: string(userHash), CountryID: 420},
		{AccountName: "CC", Email: "cc@slc.com", Password: string(userHash), CountryID: 420},
		{AccountName: "ST", Email: "st@slc.com", Password: string(userHash), CountryID: 420},
		{AccountName: "VN", Email: "vn@slc.com", Password: string(userHash), CountryID: 420},
		{AccountName: "TC", Email: "tc@slc.com", Password: string(userHash), CountryID: 420},
		{AccountName: "LL", Email: "ll@slc.com", Password: string(userHash), CountryID: 420},
		{AccountName: "GA", Email: "ga@slc.com", Password: string(userHash), CountryID: 420},
		{AccountName: "JP", Email: "jp@slc.com", Password: string(userHash), CountryID: 420},
	} {
		facades.UseDB().Clauses(clause.OnConflict{DoNothing: true}).Create(user)
	}

	return nil
}
