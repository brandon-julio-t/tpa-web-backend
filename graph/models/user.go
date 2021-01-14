package models

type User struct {
	BaseModel
	AccountName string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Password    string
	CountryID   int64
	Country     Country
}
