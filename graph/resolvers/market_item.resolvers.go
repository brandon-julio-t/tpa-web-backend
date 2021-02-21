package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"gorm.io/gorm"
)

func (r *marketItemResolver) BuyPrices(ctx context.Context, obj *models.MarketItem) ([]*models.MarketItemPrice, error) {
	rows, err := facades.UseDB().
		Model(new(models.MarketItemOffer)).
		Select("price", "count(price) as price_counts").
		Where("market_item_id = ?", obj.ID).
		Where("category = ?", "sell").
		Group("price").
		Order("price desc").
		Limit(5).
		Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	prices := make([]*models.MarketItemPrice, 0)
	for rows.Next() {
		price := new(models.MarketItemPrice)

		if err := rows.Scan(&price.Price, &price.Quantity); err != nil {
			return nil, err
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (r *marketItemResolver) Image(ctx context.Context, obj *models.MarketItem) (*models.AssetFile, error) {
	return &obj.ImageRef, facades.UseDB().Preload("ImageRef").First(obj).Error
}

func (r *marketItemResolver) PastMonthSales(ctx context.Context, obj *models.MarketItem) ([]*models.MarketItemTransaction, error) {
	now := time.Now()
	aMonthAgo := now.AddDate(0, -1, 0)

	transactions := make([]*models.MarketItemTransaction, 0)
	return transactions, facades.UseDB().
		Where("market_item_id = ?", obj.ID).
		Where("created_at between ? and ?", aMonthAgo, now).
		Find(&transactions).
		Error
}

func (r *marketItemResolver) SalePrices(ctx context.Context, obj *models.MarketItem) ([]*models.MarketItemPrice, error) {
	rows, err := facades.UseDB().
		Model(new(models.MarketItemOffer)).
		Select("price", "count(price) as price_counts").
		Where("market_item_id = ?", obj.ID).
		Where("category = ?", "buy").
		Group("price").
		Order("price desc").
		Limit(5).
		Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	prices := make([]*models.MarketItemPrice, 0)
	for rows.Next() {
		price := new(models.MarketItemPrice)

		if err := rows.Scan(&price.Price, &price.Quantity); err != nil {
			return nil, err
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (r *marketItemResolver) StartingPrice(ctx context.Context, obj *models.MarketItem) (float64, error) {
	offer := new(models.MarketItemOffer)

	if err := facades.UseDB().Where("market_item_id = ?", obj.ID).Order("price").First(offer).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}

		return 0, err
	}

	return offer.Price, nil
}

func (r *marketItemResolver) TransactionsCount(ctx context.Context, obj *models.MarketItem) (int64, error) {
	count := new(int64)
	return *count, facades.UseDB().
		Model(obj).
		Joins("join market_item_transactions mit on market_items.id = mit.market_item_id").
		Where("market_items.id = ?", obj.ID).
		Group("market_items.id").
		Count(count).
		Error
}

func (r *mutationResolver) AddMarketItemOffer(ctx context.Context, input models.AddMarketItemOffer) (*models.MarketItemOffer, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	item := new(models.MarketItem)
	if err := facades.UseDB().First(item, input.MarketItemID).Error; err != nil {
		return nil, err
	}

	matchingOffer := new(models.MarketItemOffer)
	err = facades.UseDB().
		Where("round(price, 2) = round(?, 2)", input.Price).
		First(matchingOffer).
		Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		offer := &models.MarketItemOffer{
			Category:    input.Category,
			MarketItem_: *item,
			Price:       input.Price,
			Quantity:    input.Quantity,
			User_:       *user,
		}

		message := ""
		if input.Category == "buy" {
			message = fmt.Sprintf("%v wants to buy this item for $%v", user.DisplayName, input.Price)
		} else {
			message = fmt.Sprintf("%v listed this item for sale for $%v", user.DisplayName, input.Price)
		}

		for _, socket := range r.MarketItemSockets[item.ID] {
			socket <- message
		}

		return offer, facades.UseDB().Create(offer).Error
	}

	for _, socket := range r.MarketItemSockets[item.ID] {
		if input.Category == "sell" {
			socket <- fmt.Sprintf("%v purchased this item from %v for $%v", matchingOffer.User_.DisplayName, user.DisplayName, input.Price)
		} else {
			socket <- fmt.Sprintf("%v purchased this item from %v for $%v", user.DisplayName, matchingOffer.User_.DisplayName, input.Price)
		}
	}

	err = facades.UseDB().Transaction(func(tx *gorm.DB) error {
		if input.Category == "buy" {
			if err := tx.Create(&models.MarketItemTransaction{
				Buyer_:      *user,
				Category:    "buy",
				MarketItem_: *item,
				Price:       input.Price,
				Seller_:     matchingOffer.User_,
			}).Error; err != nil {
				return err
			}

			if user.WalletBalance < input.Price {
				return errors.New("insufficient balance")
			}

			user.WalletBalance -= input.Price
			matchingOffer.User_.WalletBalance += input.Price

			if err := tx.Save(user).Error; err != nil {
				return err
			}

			if err := tx.Save(&matchingOffer.User_).Error; err != nil {
				return err
			}

			sellerInventory := new(models.Inventory)
			if err := tx.Where("user_id = ?", matchingOffer.UserID).First(sellerInventory).Error; err != nil {
				return err
			}

			if sellerInventory.Quantity == input.Quantity {
				if err := tx.Delete(sellerInventory).Error; err != nil {
					return err
				}
			} else {
				sellerInventory.Quantity -= input.Quantity
				if err := tx.Save(sellerInventory).Error; err != nil {
					return err
				}
			}

			matchingOffer.Quantity--
			if matchingOffer.Quantity == 0 {
				if err := tx.Delete(matchingOffer).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Save(matchingOffer).Error; err != nil {
					return err
				}
			}

			buyerInventory := new(models.Inventory)
			if err := facades.UseDB().
				Where("user_id = ?", user.ID).
				Where("market_item_id = ?", item.ID).
				First(buyerInventory).
				Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return tx.Create(&models.Inventory{
						User_:       *user,
						MarketItem_: *item,
						Quantity:    1,
					}).Error
				}

				buyerInventory.Quantity++
				return tx.Save(buyerInventory).Error
			}

			return tx.Delete(matchingOffer).Error
		} else {
			if err := tx.Create(&models.MarketItemTransaction{
				Buyer_:      matchingOffer.User_,
				Category:    "sell",
				MarketItem_: *item,
				Price:       input.Price,
				Seller_:     *user,
			}).Error; err != nil {
				return err
			}

			if matchingOffer.User_.WalletBalance < input.Price {
				return errors.New("insufficient balance")
			}

			user.WalletBalance += input.Price
			matchingOffer.User_.WalletBalance -= input.Price

			if err := tx.Save(user).Error; err != nil {
				return err
			}

			if err := tx.Save(&matchingOffer.User_).Error; err != nil {
				return err
			}

			sellerInventory := new(models.Inventory)
			if err := tx.Where("user_id = ?", user.ID).First(sellerInventory).Error; err != nil {
				return err
			}

			if sellerInventory.Quantity == input.Quantity {
				if err := tx.Delete(sellerInventory).Error; err != nil {
					return err
				}
			} else {
				sellerInventory.Quantity -= input.Quantity
				if err := tx.Save(sellerInventory).Error; err != nil {
					return err
				}
			}

			matchingOffer.Quantity--
			if matchingOffer.Quantity == 0 {
				if err := tx.Delete(matchingOffer).Error; err != nil {
					return err
				}
			} else {
				if err := tx.Save(matchingOffer).Error; err != nil {
					return err
				}
			}

			buyerInventory := new(models.Inventory)
			if err := facades.UseDB().
				Where("user_id = ?", matchingOffer.User_.ID).
				Where("market_item_id = ?", item.ID).
				First(buyerInventory).
				Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return tx.Create(&models.Inventory{
						User_:       matchingOffer.User_,
						MarketItem_: *item,
						Quantity:    1,
					}).Error
				}

				buyerInventory.Quantity++
				return tx.Save(buyerInventory).Error
			}

			return tx.Delete(matchingOffer).Error
		}
	})

	now := time.Now()
	if input.Category == "buy" {
		buyerMessage := fmt.Sprintf("You purchased item %v for $%v at %v.", item.Name, input.Price, now)
		sellerMessage := fmt.Sprintf("%v purchased item %v from you for $%v at %v.", user.DisplayName, item.Name, input.Price, now)

		if err := facades.UseMail().SendText(buyerMessage, "Market Item Purchase", "MarketItemPurchase", user.Email); err != nil {
			return nil, err
		}

		if err := facades.UseMail().SendText(sellerMessage, "Market Item Purchase", "MarketItemPurchase", matchingOffer.User_.Email); err != nil {
			return nil, err
		}
	} else {
		buyerMessage := fmt.Sprintf("You purchased item %v for $%v at %v.", item.Name, input.Price, now)
		sellerMessage := fmt.Sprintf("%v purchased item %v from you for $%v at %v.", matchingOffer.User_.DisplayName, item.Name, input.Price, now)

		if err := facades.UseMail().SendText(buyerMessage, "Market Item Purchase", "MarketItemPurchase", matchingOffer.User_.Email); err != nil {
			return nil, err
		}

		if err := facades.UseMail().SendText(sellerMessage, "Market Item Purchase", "MarketItemPurchase", user.Email); err != nil {
			return nil, err
		}
	}

	return matchingOffer, err
}

func (r *mutationResolver) CancelMarketItemOffer(ctx context.Context, id int64) (*models.MarketItemOffer, error) {
	offer := new(models.MarketItemOffer)

	if err := facades.UseDB().First(offer, id).Error; err != nil {
		return nil, err
	}

	return offer, facades.UseDB().Delete(offer).Error
}

func (r *queryResolver) MarketItem(ctx context.Context, id int64) (*models.MarketItem, error) {
	item := new(models.MarketItem)
	return item, facades.UseDB().First(item, id).Error
}

func (r *queryResolver) MarketItems(ctx context.Context, page int64) (*models.MarketItemPagination, error) {
	items := make([]*models.MarketItem, 0)
	count := new(int64)
	perPage := 10

	if err := facades.UseDB().
		Model(new(models.MarketItem)).
		Joins("full join market_item_transactions mit on market_items.id = mit.market_item_id").
		Group("market_items.id").
		Count(count).
		Scopes(facades.UsePagination(int(page), perPage)).
		Order("count(mit.id) desc").
		Find(&items).
		Error; err != nil {
		return nil, err
	}

	return &models.MarketItemPagination{
		Data:       items,
		TotalPages: int64(math.Ceil(float64(*count) / float64(perPage))),
	}, nil
}

func (r *subscriptionResolver) OnMarketItemOfferAdded(ctx context.Context, marketItemID int64) (<-chan string, error) {
	socket := make(chan string, 1)
	r.MarketItemSockets[marketItemID] = append(r.MarketItemSockets[marketItemID], socket)
	return socket, nil
}

func (r *userResolver) GamesByOwnedMarketItems(ctx context.Context, obj *models.User) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	if err := facades.UseDB().
		Joins("join market_items mi on games.id = mi.game_id").
		Joins("join inventories i on mi.id = i.market_item_id").
		Where("user_id = ?", obj.ID).
		Distinct("games.*").
		Find(&games).
		Error; err != nil {
		return nil, err
	}

	return games, nil
}

func (r *userResolver) MarketItemsBuyListing(ctx context.Context, obj *models.User) ([]*models.MarketItemOffer, error) {
	offers := make([]*models.MarketItemOffer, 0)
	return offers, facades.UseDB().
		Preload("MarketItem_.Game_").
		Where("user_id = ?", obj.ID).
		Where("category = ?", "buy").
		Find(&offers).
		Error
}

func (r *userResolver) MarketItemsByGame(ctx context.Context, obj *models.User, page int64, gameID int64, filter string) (*models.MarketItemPagination, error) {
	items := make([]*models.MarketItem, 0)
	count := new(int64)
	perPage := 10

	inventories := make([]*models.Inventory, 0)
	if err := facades.UseDB().
		Model(new(models.Inventory)).
		Joins("join market_items mi on inventories.market_item_id = mi.id").
		Where("game_id = ?", gameID).
		Where("user_id = ?", obj.ID).
		Where("(lower(mi.category) like lower(?) or lower(mi.name) like lower(?))", "%"+filter+"%", "%"+filter+"%").
		Count(count).
		Scopes(facades.UsePagination(int(page), perPage)).
		Find(&inventories).
		Error; err != nil {
		return nil, err
	}

	for _, inventory := range inventories {
		if err := facades.UseDB().
			Preload("MarketItem_").
			Preload("MarketItem_.Game_").
			First(inventory).
			Error; err != nil {
			return nil, err
		}

		items = append(items, &inventory.MarketItem_)
	}

	return &models.MarketItemPagination{
		Data:       items,
		TotalPages: int64(math.Ceil(float64(*count) / float64(perPage))),
	}, nil
}

func (r *userResolver) MarketItemsSellListing(ctx context.Context, obj *models.User) ([]*models.MarketItemOffer, error) {
	offers := make([]*models.MarketItemOffer, 0)
	return offers, facades.UseDB().
		Preload("MarketItem_.Game_").
		Where("user_id = ?", obj.ID).
		Where("category = ?", "sell").
		Find(&offers).
		Error
}

// MarketItem returns generated.MarketItemResolver implementation.
func (r *Resolver) MarketItem() generated.MarketItemResolver { return &marketItemResolver{r} }

type marketItemResolver struct{ *Resolver }
