package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func SeedGameGenres() error {
	return facades.UseDB().Create([]*models.GameGenre{
		{ID: 1, Name: "Action"},
		{ID: 2, Name: "Adventure"},
		{ID: 3, Name: "Casual"},
		{ID: 4, Name: "Early Access"},
		{ID: 5, Name: "Free to Play"},
		{ID: 6, Name: "Indie"},
		{ID: 7, Name: "Massively Multiplayer"},
		{ID: 8, Name: "Racing"},
		{ID: 9, Name: "RPG"},
		{ID: 10, Name: "Simulation"},
		{ID: 11, Name: "Sports"},
		{ID: 12, Name: "Strategy"},
	}).Error
}
