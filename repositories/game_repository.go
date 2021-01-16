package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"io/ioutil"
)

type GameRepository struct{}

func (GameRepository) GetAll(page int) ([]*models.Game, error) {
	var games []*models.Game
	return games, facades.UseDB().Scopes(facades.UsePagination(page, 3)).Preload("Slideshows").Preload("GameTags").Find(&games).Error
}

func (GameRepository) Create(input models.CreateGame) (*models.Game, error) {
	banner, err := ioutil.ReadAll(input.Banner.File)
	if err != nil {
		return nil, err
	}

	var slideshows []*models.GameSlideshow

	for _, slideshowInput := range input.Slideshows {
		data, err := ioutil.ReadAll(slideshowInput.File)
		if err != nil {
			return nil, err
		}

		slideshows = append(slideshows, &models.GameSlideshow{
			GameID: 0,
			Game:   models.Game{},
			File:   data,
		})
	}

	var gameTags []*models.GameTag

	for _, gameTagId := range input.GameTags {
		var gameTag models.GameTag

		if err := facades.UseDB().First(&gameTag, gameTagId).Error; err != nil {
			return nil, err
		}

		gameTags = append(gameTags, &gameTag)
	}

	game := models.Game{
		Title:              input.Title,
		Description:        input.Description,
		Price:              input.Price,
		Banner:             banner,
		Slideshows:         slideshows,
		GameTags:           gameTags,
		SystemRequirements: input.SystemRequirements,
	}

	if err := facades.UseDB().Create(&game).Error; err != nil {
		return nil, err
	}

	return &game, nil
}

func (GameRepository) Delete(id int64) (*models.Game, error)  {
	var game models.Game
	return &game, facades.UseDB().Delete(&game, id).Error
}
