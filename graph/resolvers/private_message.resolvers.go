package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"github.com/brandon-julio-t/tpa-web-backend/services/notification_service"
)

func (r *mutationResolver) AddPrivateMessage(ctx context.Context, friendID int64, text string) (*models.PrivateMessage, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	var friend models.User
	if err := facades.UseDB().First(&friend, friendID).Error; err != nil {
		return nil, err
	}

	privateMessage := models.PrivateMessage{
		Text:     text,
		Sender:   *user,
		Receiver: friend,
	}

	if err := facades.UseDB().Create(&privateMessage).Error; err != nil {
		return nil, err
	}

	if friendSocket, ok := r.PrivateChatSockets[friendID]; ok {
		friendSocket <- &privateMessage
	} else {
		if err := notification_service.Notify(&friend, fmt.Sprintf("%v sent you a message", user.DisplayName)); err != nil {
			return nil, err
		}
	}

	return &privateMessage, nil
}

func (r *queryResolver) PrivateMessage(ctx context.Context, friendID int64) ([]*models.PrivateMessage, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	var messages []*models.PrivateMessage

	return messages, facades.UseDB().
		Where("sender_id = ? or receiver_id = ?", friendID, friendID).
		Where("sender_id = ? or receiver_id = ?", user.ID, user.ID).
		Order("created_at desc").
		Preload("Sender").
		Preload("Receiver").
		Find(&messages).
		Error
}

func (r *subscriptionResolver) PrivateMessageAdded(ctx context.Context) (<-chan *models.PrivateMessage, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	socket := make(chan *models.PrivateMessage, 1)

	r.PrivateChatSockets[user.ID] = socket

	go func() {
		<-ctx.Done()
		delete(r.PrivateChatSockets, user.ID)
	}()

	return socket, nil
}

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type subscriptionResolver struct{ *Resolver }
