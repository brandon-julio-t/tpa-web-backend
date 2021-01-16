package facades

import (
	"fmt"
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

}

func UseDB() *gorm.DB {
	return databaseSingleton
}
