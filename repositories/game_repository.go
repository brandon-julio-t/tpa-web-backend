package repositories

import (
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"gorm.io/gorm"
	"io/ioutil"
	"math"
	"time"
)

type GameRepository struct{}

func (GameRepository) GetAll(page int) (*models.GamePagination, error) {
	var games []*models.Game
	var totalGames int64
	if err := usePreloadedGame().
		Model(&models.Game{}).
		Count(&totalGames).
		Scopes(facades.UsePagination(page, 3)).
		Find(&games).
		Error; err != nil {
		return nil, err
	}
	return &models.GamePagination{
		Data:       games,
		TotalPages: int64(math.Ceil(float64(totalGames) / float64(3))),
	}, nil
}

func usePreloadedGame() *gorm.DB {
	return facades.UseDB().
		Preload("Banner").
		Preload("GameTags").
		Preload("Genre").
		Preload("GameSlideshows").
		Preload("GameSlideshows.GameSlideshowFile")
}

func (GameRepository) GetFeaturedAndRecommendedGames() ([]*models.Game, error) {
	var games []*models.Game
	return games, usePreloadedGame().
		Order("hours_played desc").
		Limit(5).
		Find(&games).
		Error
}

func (GameRepository) GetById(id int64) (*models.Game, error) {
	var game models.Game
	return &game, usePreloadedGame().First(&game, id).Error
}

func (GameRepository) GetCommunityRecommends() *gorm.DB {
	now := time.Now()
	aWeekAgo := now.AddDate(0, 0, -7)

	return facades.UseDB().
		Joins(
			"join (?) as recommended_games on games.id = recommended_games.id",
			facades.UseDB().
				Select("g.id, count(game_reviews.id) as recommendations_count").
				Model(new(models.GameReview)).
				Joins("join games g on game_reviews.game_review_game_id = g.id").
				Where("is_recommended = ?", true).
				Where("game_reviews.created_at >= ?", aWeekAgo).
				Where("game_reviews.created_at <= ?", now).
				Group("g.id"),
		).
		Joins(
			"join (?) as unrecommended_games on games.id = unrecommended_games.id",
			facades.UseDB().
				Select("g.id, count(game_reviews.id) as recommendations_count").
				Model(new(models.GameReview)).
				Joins("join games g on game_reviews.game_review_game_id = g.id").
				Where("is_recommended = ?", false).
				Where("game_reviews.created_at >= ?", aWeekAgo).
				Where("game_reviews.created_at <= ?", now).
				Group("g.id"),
		).
		Order("recommended_games.recommendations_count - unrecommended_games.recommendations_count desc")
}

func (GameRepository) GetSpecialOffers() *gorm.DB {
	return facades.UseDB().
		Where("discount > 0").
		Order("discount desc")
}

func (GameRepository) GetTopSellers() *gorm.DB {
	aWeekAgo := time.Now().AddDate(0, 0, -7)
	return facades.UseDB().
		Joins(
			"join (?) as purchases on games.id = purchases.game_id",
			facades.UseDB().
				Model(new(models.GamePurchaseTransactionHeader)).
				Select("gptd.game_purchase_transaction_detail_game_id as game_id, count(gptd.game_purchase_transaction_header_id) as total_purchases").
				Joins("join game_purchase_transaction_details gptd on game_purchase_transaction_headers.id = gptd.game_purchase_transaction_header_id").
				Where("game_purchase_transaction_headers.created_at >= ?", aWeekAgo).
				Group("gptd.game_purchase_transaction_detail_game_id"),
		).
		Joins(
			"join (?) as gifts on games.id = gifts.game_id",
			facades.UseDB().
				Model(new(models.GameGiftTransactionHeader)).
				Select("ggtd.game_gift_transaction_detail_game_id as game_id, count(ggtd.game_gift_transaction_header_id) as total_gifts").
				Joins("join game_gift_transaction_details ggtd on game_gift_transaction_headers.id = ggtd.game_gift_transaction_header_id").
				Where("game_gift_transaction_headers.created_at >= ?", aWeekAgo).
				Group("ggtd.game_gift_transaction_detail_game_id"),
		).
		Order("total_purchases + total_gifts desc")
}

func (GameRepository) Create(input models.CreateGame) (*models.Game, error) {
	var game models.Game

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
			GameSlideshowFile: models.AssetFile{
				File:        data,
				ContentType: slideshowInput.ContentType,
			},
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

	gameBanner := models.AssetFile{File: banner, ContentType: input.Banner.ContentType}
	if err := facades.UseDB().Create(&gameBanner).Error; err != nil {
		return nil, err
	}

	game = models.Game{
		Title:              input.Title,
		Description:        input.Description,
		Price:              input.Price,
		BannerID:           gameBanner.ID,
		Banner:             gameBanner,
		GenreID:            input.GenreID,
		IsInappropriate:    input.IsInappropriate,
		GameSlideshows:     slideshows,
		GameTags:           gameTags,
		SystemRequirements: input.SystemRequirements,
	}

	return &game, facades.UseDB().Create(&game).Error
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
	game.IsInappropriate = input.IsInappropriate

	var genre models.GameGenre

	if err := facades.UseDB().First(&genre, input.GenreID).Error; err != nil {
		return nil, err
	}

	game.Genre = genre
	game.GenreID = input.GenreID

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

		if err := tx.Delete(game.GameSlideshows).Error; err != nil {
			return err
		}

		for _, slideshow := range game.GameSlideshows {
			if err := tx.Delete(&slideshow.GameSlideshowFile).Error; err != nil {
				return err
			}
		}

		var slideshows []*models.GameSlideshow

		for _, slideshow := range input.Slideshows {
			slideshowData, err := ioutil.ReadAll(slideshow.File)
			if err != nil {
				return err
			}

			slideshows = append(slideshows, &models.GameSlideshow{GameSlideshowFile: models.AssetFile{
				File:        slideshowData,
				ContentType: slideshow.ContentType,
			}})
		}

		game.GameSlideshows = slideshows

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

		if err := tx.Delete(game.GameSlideshows).Error; err != nil {
			return err
		}

		for _, slideshow := range game.GameSlideshows {
			if err := tx.Delete(slideshow.GameSlideshowFile).Error; err != nil {
				return err
			}
		}

		if err := tx.Delete(&game).Error; err != nil {
			return err
		}

		return tx.Delete(&game.Banner).Error
	})
}
