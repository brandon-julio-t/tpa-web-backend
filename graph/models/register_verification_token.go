package models

type RegisterVerificationToken struct {
	ID    int64  `gorm:"primaryKey"`
	Token string `gorm:"uniqueIndex"`
}
