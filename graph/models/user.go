package models

import (
	"encoding/base64"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io/ioutil"
	"path/filepath"
	"time"
)

type User struct {
	ID             int64  `gorm:"primaryKey"`
	AccountName    string `gorm:"uniqueIndex"`
	Country        Country
	CountryID      int64
	CustomURL      string `gorm:"uniqueIndex"`
	DisplayName    string
	Email          string `gorm:"uniqueIndex"`
	Password       string
	ProfilePicture []byte
	ProfileTheme   string
	RealName       string
	Summary        string
	WalletBalance  float64
	SuspendedAt    time.Time `gorm:"index"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	path := filepath.Join("assets", "default_profile_picture.png")
	defaultProfilePicture, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	u.ProfileTheme = "#4B5563"
	u.ProfilePicture = defaultProfilePicture
	u.DisplayName = u.AccountName
	u.CustomURL = uuid.Must(uuid.NewRandom()).String()

	return nil
}

func (u *User) ProfilePictureBase64() string {
	return base64.StdEncoding.EncodeToString(u.ProfilePicture)
}

func (u *User) ReportCounts() int64 {
	var count int64
	if err := facades.UseDB().Model(&Report{}).Where("reported_id = ?", u.ID).Count(&count).Error; err != nil {
		return -1
	}
	return count
}
