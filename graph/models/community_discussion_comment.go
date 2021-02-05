package models

import "time"

type CommunityDiscussionComment struct {
	ID                    int64
	Body                  string
	CommunityDiscussionID int64
	CreatedAt             time.Time
	UserID                int64

	CommunityDiscussion_ CommunityDiscussion `gorm:"foreignKey:CommunityDiscussionID"`
	User_                User                `gorm:"foreignKey:UserID"`
}
