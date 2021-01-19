package database_seeds

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"io/ioutil"
	"path/filepath"
)

func SeedGames() error {
	path := filepath.Join("assets", "Background Zoom SLC.png")
	backgroundSLC, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	return facades.UseDB().Create(&[]*models.Game{
		{
			Title:       "The Butcher 3",
			Description: "An old man living with his sword hunting monsters",
			Price:       499,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Grand Theft Manual",
			Description: "Player and the world",
			Price:       499,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Call of Sleeping",
			Description: "The most popular FPS game",
			Price:       299,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Counter Stroke: Local Offensive",
			Description: "An old man living with his sword hunting monsters",
			Price:       99,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Sleeping Field",
			Description: "Another most popular FPS game",
			Price:       399,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Sleeping Cats",
			Description: "A man fighting thugs in China",
			Price:       399,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Gacha Impact",
			Description: "Fantasy anime game",
			Price:       199,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Cybercafe 2069",
			Description: "What makes someone a criminal in 2069?",
			Price:       499,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "The Last of Them",
			Description: "The civilization in post apocalypse world settings",
			Price:       499,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
		{
			Title:       "Potato Gear Rising",
			Description: "Rules of potato",
			Price:       499,
			Banner:      models.AssetFile{File: backgroundSLC, ContentType: "image/png"},
			Slideshows: []*models.GameSlideshow{
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
				{File: models.AssetFile{File: backgroundSLC, ContentType: "image/png"}},
			},
			GameTags:           []*models.GameTag{{ID: 1}, {ID: 2}, {ID: 3}},
			SystemRequirements: "Potato PC",
		},
	}).Error
}
