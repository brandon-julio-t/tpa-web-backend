package facades

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var databaseSingleton *gorm.DB = nil

func init() {
	if dbUrl := os.Getenv("DATABASE_URL"); dbUrl != "" {
		db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				logger.Config{
					SlowThreshold: time.Second,
					LogLevel:      logger.Info,
					Colorful:      true,
				},
			),
		})
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
