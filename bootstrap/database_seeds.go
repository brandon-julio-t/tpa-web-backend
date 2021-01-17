package bootstrap

import (
	"github.com/brandon-julio-t/tpa-web-backend/database_seeds"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"gorm.io/gorm"
	"log"
)

func init() {
	if err := facades.UseDB().Transaction(func(tx *gorm.DB) error {
		for _, seed := range []func() error{
			database_seeds.SeedCountries,
			database_seeds.SeedUsers,
			database_seeds.SeedGameTags,
		} {
			if err := seed(); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}
}
