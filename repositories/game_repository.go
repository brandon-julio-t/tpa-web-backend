package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"io/ioutil"
	"log"
)

type GameRepository struct{}

func (GameRepository) GetAll(page int) ([]*models.Game, error) {
	var games []*models.Game
	return games, facades.UseDB().Scopes(facades.UsePagination(page, 3)).Preload("Slideshows").Preload("GameTags").Find(&games).Error
}

func (GameRepository) GetById(id int64) (*models.Game, error) {
	var game models.Game
	return &game, facades.UseDB().Preload("Slideshows").Preload("GameTags").First(&game, id).Error
}

func (GameRepository) Create(input models.CreateGame) (*models.Game, error) {
	banner, err := ioutil.ReadAll(input.Banner.File)
	if err != nil {
		return nil, err
	}

	var slideshows []*models.GameSlideshow

	for _, slideshowInput := range input.Slideshows {
		log.Print(slideshowInput.ContentType)

		data, err := ioutil.ReadAll(slideshowInput.File)
		if err != nil {
			return nil, err
		}

		slideshows = append(slideshows, &models.GameSlideshow{
			File:        data,
			ContentType: slideshowInput.ContentType,
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

func (GameRepository) Delete(id int64) (*models.Game, error) {
	var game models.Game

	if err := facades.UseDB().Preload("Slideshows").Preload("GameTags").First(&game, id).Error; err != nil {
		return nil, err
	}

	if err := facades.UseDB().Model(&game).Association("GameTags").Clear(); err != nil {
		return nil, err
	}

	if err := facades.UseDB().Delete(game.Slideshows).Error; err != nil {
		return nil, err
	}

	return &game, facades.UseDB().Delete(&game).Error
}
