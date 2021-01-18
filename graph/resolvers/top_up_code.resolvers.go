package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"gorm.io/gorm"
)

func (r *mutationResolver) RedeemWallet(ctx context.Context, code string) (bool, error) {
	user := middlewares.UseAuth(ctx)
	if user == nil {
		return false, errors.New("not authenticated")
	}

	return true, facades.UseDB().Transaction(func(tx *gorm.DB) error {
		var topUpCode models.TopUpCode

		if err := tx.First(&topUpCode, "code = ?", code).Error; err != nil {
			return err
		}

		user.WalletBalance += topUpCode.Amount

		if err := tx.Delete(&topUpCode).Error; err != nil {
			return err
		}

		return tx.Save(&user).Error
	})
}
