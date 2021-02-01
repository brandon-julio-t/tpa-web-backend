package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
)

func (r *gameReviewResolver) Comment(ctx context.Context, obj *models.GameReview, page int64) (*models.GameReviewCommentPagination, error) {
	comments := make([]*models.GameReviewComment, 0)
	count := new(int64)
	perPage := 10

	if err := facades.UseDB().
		Model(new(models.GameReviewComment)).
		Where("game_review_id = ?", obj.ID).
		Count(count).
		Scopes(facades.UsePagination(int(page), perPage)).
		Find(&comments).
		Error; err != nil {
		return nil, err
	}

	return &models.GameReviewCommentPagination{
		Data:       comments,
		TotalPages: int64(math.Ceil(float64(*count) / float64(perPage))),
	}, nil
}

func (r *mutationResolver) PostGameReviewComment(ctx context.Context, input models.GameReviewCommentInput) (*models.GameReviewComment, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	review := new(models.GameReview)
	if err := facades.UseDB().First(review, input.ReviewID).Error; err != nil {
		return nil, err
	}

	comment := &models.GameReviewComment{
		Body:  input.Body,
		GameReview_: *review,
		User_: *user,
	}
	return comment, facades.UseDB().Create(comment).Error
}
