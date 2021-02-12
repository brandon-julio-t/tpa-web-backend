package models

import (
	"context"
	"fmt"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"syreclabs.com/go/faker"
	"time"
)

type User struct {
	ID                      int64  `gorm:"primaryKey"`
	AccountName             string `gorm:"uniqueIndex"`
	Country                 Country
	CountryID               int64
	CustomURL               string `gorm:"uniqueIndex"`
	DisplayName             string
	Email                   string `gorm:"uniqueIndex"`
	Exp                     int64
	Friends                 []*Friendship
	FriendCode              string        `gorm:"uniqueIndex"`
	MiniProfileBackgroundID int64
	Password                string
	UserProfilePictureID    int64
	UserProfilePicture      AssetFile
	Points                  int64
	ProfileBackgroundID     int64
	AvatarBorderID          int64
	ProfileTheme            string
	RealName                string
	Status                  string
	Summary                 string
	WalletBalance           float64
	UserWishlist            []*Game   `gorm:"many2many:wishlists;"`
	UserCart                []*Game   `gorm:"many2many:cart;"`
	SuspendedAt             time.Time `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	path := filepath.Join("assets", "default_profile_picture.png")
	defaultProfilePicture, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	u.CustomURL = uuid.Must(uuid.NewRandom()).String()
	u.DisplayName = u.AccountName
	u.FriendCode = faker.Numerify("#########")
	u.ProfileTheme = "#4B5563"
	u.UserProfilePicture = AssetFile{File: defaultProfilePicture, ContentType: "image/png"}

	if u.Status == "" {
		u.Status = "offline"
	}

	return nil
}

func (u *User) ReportCounts(ctx context.Context) int64 {
	cacheKey := fmt.Sprintf("users:%v:reports_count", u.ID)
	cached, err := facades.UseCache().Get(ctx, cacheKey).Result()
	if err != redis.Nil && cached != "" {
		result, err := strconv.ParseInt(cached, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		return result
	}

	var count int64
	if err := facades.UseDB().Model(&Report{}).Where("reported_id = ?", u.ID).Count(&count).Error; err != nil {
		return -1
	}
	if err := facades.UseCache().Set(ctx, cacheKey, count, 10*time.Second).Err(); err != nil {
		log.Fatal(err)
	}
	return count
}
