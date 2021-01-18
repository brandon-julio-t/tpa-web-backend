package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"io/ioutil"
)

type GameRepository struct{}

func (GameRepository) GetAll(page int) ([]*models.Game, error) {
	var games []*models.Game
	return games, usePreloadedGame().Scopes(facades.UsePagination(page, 3)).Find(&games).Error
}

func usePreloadedGame() *gorm.DB {
	return facades.UseDB().
		Preload("Banner").
		Preload("Slideshows").
		Preload("Slideshows.File").
		Preload("GameTags")
}

func (GameRepository) GetById(id int64) (*models.Game, error) {
	var game models.Game
	return &game, usePreloadedGame().First(&game, id).Error
}

func (GameRepository) Create(input models.CreateGame) (*models.Game, error) {
	var game models.Game
	return &game, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		banner, err := ioutil.ReadAll(input.Banner.File)
		if err != nil {
			return err
		}

		var slideshows []*models.GameSlideshow

		for _, slideshowInput := range input.Slideshows {
			data, err := ioutil.ReadAll(slideshowInput.File)
			if err != nil {
				return err
			}

			slideshows = append(slideshows, &models.GameSlideshow{
				File: models.AssetFile{
					File:        data,
					ContentType: slideshowInput.ContentType,
				},
			})
		}

		var gameTags []*models.GameTag

		for _, gameTagId := range input.GameTags {
			var gameTag models.GameTag

			if err := tx.First(&gameTag, gameTagId).Error; err != nil {
				return err
			}

			gameTags = append(gameTags, &gameTag)
		}

		game = models.Game{
			Title:              input.Title,
			Description:        input.Description,
			Price:              input.Price,
			Banner:             models.AssetFile{File: banner, ContentType: input.Banner.ContentType},
			Slideshows:         slideshows,
			GameTags:           gameTags,
			SystemRequirements: input.SystemRequirements,
		}

		return tx.Create(&game).Error
	})
}

func (r GameRepository) Update(input models.UpdateGame) (*models.Game, error) {
	var game models.Game

	if err := usePreloadedGame().First(&game, input.ID).Error; err != nil {
		return nil, err
	}

	game.SystemRequirements = input.SystemRequirements
	game.Price = input.Price
	game.Title = input.Title
	game.Description = input.Description

	return &game, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&game).Association("GameTags").Clear(); err != nil {
			return err
		}

		var tags []*models.GameTag
		for _, tag := range input.GameTags {
			var gameTag models.GameTag
			if err := tx.First(&gameTag, tag).Error; err != nil {
				return err
			}
			tags = append(tags, &gameTag)
		}

		game.GameTags = tags

		bannerData, err := ioutil.ReadAll(input.Banner.File)
		if err != nil {
			return err
		}

		game.Banner.File = bannerData
		game.Banner.ContentType = input.Banner.ContentType

		if err := tx.Save(&game.Banner).Error; err != nil {
			return err
		}

		if err := tx.Delete(game.Slideshows).Error; err != nil {
			return err
		}

		for _, slideshow := range game.Slideshows {
			if err := tx.Delete(&slideshow.File).Error; err != nil {
				return err
			}
		}

		var slideshows []*models.GameSlideshow

		for _, slideshow := range input.Slideshows {
			slideshowData, err := ioutil.ReadAll(slideshow.File)
			if err != nil {
				return err
			}

			slideshows = append(slideshows, &models.GameSlideshow{File: models.AssetFile{
				File:        slideshowData,
				ContentType: slideshow.ContentType,
			}})
		}

		game.Slideshows = slideshows

		return tx.Save(&game).Error
	})
}

func (GameRepository) Delete(id int64) (*models.Game, error) {
	var game models.Game

	if err := usePreloadedGame().First(&game, id).Error; err != nil {
		return nil, err
	}

	return &game, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&game).Association("GameTags").Clear(); err != nil {
			return err
		}

		if err := tx.Delete(game.Slideshows).Error; err != nil {
			return err
		}

		for _, slideshow := range game.Slideshows {
			if err := tx.Delete(slideshow.File).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(&game).Error; err != nil {
			return err
		}

		return tx.Delete(&game.Banner).Error
	})
}
