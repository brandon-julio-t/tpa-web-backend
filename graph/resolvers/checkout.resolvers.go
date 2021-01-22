package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/brandon-julio-t/tpa-web-backend/commands"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"gorm.io/gorm"
)

func (r *mutationResolver) CheckoutWithWallet(ctx context.Context) (float64, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return 0, err
	}

	var games []*models.Game
	if err := facades.UseDB().Model(user).Association("UserCart").Find(&games); err != nil {
		return 0, err
	}

	var grandTotal float64
	for _, game := range games {
		grandTotal += game.Price
	}

	if user.WalletBalance < grandTotal {
		return 0, errors.New("insufficient balance")
	}

	return grandTotal, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		user.WalletBalance -= grandTotal

		if err := tx.Save(user).Error; err != nil {
			return err
		}

		return commands.GamePurchaseCommand{
			DB:         tx,
			Games:      games,
			User:       *user,
			GrandTotal: grandTotal,
		}.Execute()
	})
}

func (r *mutationResolver) CheckoutWithCard(ctx context.Context) (float64, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return 0, err
	}

	var games []*models.Game
	if err := facades.UseDB().Model(user).Association("UserCart").Find(&games); err != nil {
		return 0, err
	}

	var grandTotal float64
	for _, game := range games {
		grandTotal += game.Price
	}

	return grandTotal, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		return commands.GamePurchaseCommand{
			DB:         tx,
			Games:      games,
			User:       *user,
			GrandTotal: grandTotal,
		}.Execute()
	})
}

func (r *mutationResolver) GiftWithWallet(ctx context.Context, input models.Gift) (float64, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return 0, err
	}

	var games []*models.Game
	if err := facades.UseDB().Model(user).Association("UserCart").Find(&games); err != nil {
		return 0, err
	}

	var grandTotal float64
	for _, game := range games {
		grandTotal += game.Price
	}

	if user.WalletBalance < grandTotal {
		return 0, errors.New("insufficient balance")
	}

	return grandTotal, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		user.WalletBalance -= grandTotal

		if err := tx.Save(user).Error; err != nil {
			return err
		}

		return commands.GameGiftCommand{
			DB:    tx,
			User:  *user,
			Games: games,
			Input: input,
		}.Execute()
	})
}

func (r *mutationResolver) GiftWithCard(ctx context.Context, input models.Gift) (float64, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return 0, err
	}

	var games []*models.Game
	if err := facades.UseDB().Model(user).Association("UserCart").Find(&games); err != nil {
		return 0, err
	}

	var grandTotal float64
	for _, game := range games {
		grandTotal += game.Price
	}

	return grandTotal, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		return commands.GameGiftCommand{
			DB:    tx,
			User:  *user,
			Games: games,
			Input: input,
		}.Execute()
	})
}
