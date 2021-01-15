package models

import (
	"encoding/base64"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io/ioutil"
	"path/filepath"
)

type User struct {
	BaseModel
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
