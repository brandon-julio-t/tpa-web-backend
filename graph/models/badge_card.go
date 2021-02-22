package models

type BadgeCard struct {
	ID      int64
	BadgeID int64
	Badge   Badge
	Name    string
}
