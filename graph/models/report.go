package models

type Report struct {
	BaseModel
	ReporterID  int64 `gorm:"primaryKey"`
	Reporter    User
	ReportedID  int64 `gorm:"primaryKey"`
	Reported    User
	Description string
}
