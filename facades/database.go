package facades

import (
	"fmt"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var databaseSingleton *gorm.DB = nil

func init() {
	if dbUrl, ok := os.LookupEnv("DATABASE_URL"); ok {
		db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		databaseSingleton = db
	} else {
		dsn := fmt.Sprintf(
			"host=%v user=%v password=%v dbname=%v port=%v",
			"localhost",
			"postgres",
			"postgres",
			"tpa-web",
			"5432",
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		databaseSingleton = db
	}

	runMigration()
	runSeed()
}

func runMigration() {
	if err := databaseSingleton.AutoMigrate(
		&models.User{},
		&models.Country{},
		&models.RegisterVerificationToken{},
	); err != nil {
		log.Fatal(err)
	}
}

func runSeed() {
	var country models.Country
	if err := UseDB().First(&country, "name = ?", "Indonesia").Error; err != nil {
		log.Fatal(err)
	}

	adminHash, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	if err := UseDB().First(&models.User{}, "email = ", "admin@admin.com").Error; err != nil {
		UseDB().Create(&models.User{
			AccountName: "Admin",
			Email:       "admin@admin.com",
			Password:    string(adminHash),
			CountryID:   country.ID,
			Country:     country,
		})
	}

	userHash, err := bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	if err := UseDB().First(&models.User{}, "email = ", "user@user.com").Error; err != nil {
		UseDB().Create(&models.User{
			AccountName: "User",
			Email:       "user@user.com",
			Password:    string(userHash),
			CountryID:   country.ID,
			Country:     country,
		})
	}
}

func UseDB() *gorm.DB {
	return databaseSingleton
}
