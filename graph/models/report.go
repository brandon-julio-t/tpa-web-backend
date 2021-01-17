package models

import "time"

type Report struct {
	ID          int64 `gorm:"primaryKey"`
	CreatedAt   time.Time
	ReporterID  int64 `gorm:"primaryKey"`
	Reporter    User
	ReportedID  int64 `gorm:"primaryKey"`
	Reported    User
	Description string
}
