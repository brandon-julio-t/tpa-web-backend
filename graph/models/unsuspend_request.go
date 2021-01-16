package models

type UnsuspendRequest struct {
	BaseModel
	UserID int64
	User   User
}
