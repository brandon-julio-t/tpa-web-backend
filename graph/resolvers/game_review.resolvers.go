package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
	"gorm.io/gorm"
)

func (r *gameResolver) MostHelpfulReviews(ctx context.Context, obj *models.Game) ([]*models.GameReview, error) {
	reviews := make([]*models.GameReview, 0)
	now := time.Now()
	thirtyDaysAgo := now.AddDate(0, 0, -30)

	rows, err := facades.UseDB().
		Raw(`
select is_recommended, content, gr.id, game_review_user_id, g.created_at, game_review_game_id
from games g
         join game_reviews gr on g.id = gr.game_review_game_id
         join (select game_review_vote_game_review_id, count(*) as value
               from game_review_votes
               where is_up_vote = true
               group by game_review_vote_game_review_id) as positive_votes
              on gr.id = positive_votes.game_review_vote_game_review_id
         join (select game_review_vote_game_review_id, count(*) as value
               from game_review_votes
               where is_up_vote = false
               group by game_review_vote_game_review_id) as negative_votes
              on gr.id = negative_votes.game_review_vote_game_review_id
where gr.created_at >= ?
  and gr.created_at <= ?
  and g.id = ?
order by positive_votes.value - negative_votes.value desc
`, thirtyDaysAgo, now, obj.ID).Rows()

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		review := new(models.GameReview)
		if err := rows.Scan(
			&review.IsRecommended,
			&review.Content,
			&review.ID,
			&review.GameReviewUserID,
			&review.CreatedAt,
			&review.GameReviewGameID,
		); err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}

	return reviews, rows.Close()
}

func (r *gameResolver) RecentReviews(ctx context.Context, obj *models.Game) ([]*models.GameReview, error) {
	return obj.GameGameReviews, facades.UseDB().
		Preload("GameGameReviews").
		Order("created_at desc").
		First(obj).
		Error
}

func (r *gameReviewResolver) User(ctx context.Context, obj *models.GameReview) (*models.User, error) {
	return &obj.GameReviewUser, facades.UseDB().Preload("GameReviewUser").First(obj).Error
}

func (r *gameReviewResolver) UpVotes(ctx context.Context, obj *models.GameReview) (int64, error) {
	count := int64(0)
	return count, facades.UseDB().
		First(obj).
		Joins("join game_review_votes grv on game_reviews.id = grv.game_review_vote_game_review_id").
		Where("is_up_vote = ?", true).
		Count(&count).
		Error
}

func (r *gameReviewResolver) DownVotes(ctx context.Context, obj *models.GameReview) (int64, error) {
	count := int64(0)
	return count, facades.UseDB().
		First(obj).
		Joins("join game_review_votes grv on game_reviews.id = grv.game_review_vote_game_review_id").
		Where("is_up_vote = ?", false).
		Count(&count).
		Error
}

func (r *gameReviewResolver) UpVoters(ctx context.Context, obj *models.GameReview) ([]*models.User, error) {
	users := make([]*models.User, 0)

	if err := facades.UseDB().
		Joins("join game_review_votes grv on users.id = grv.game_review_vote_user_id").
		Where("is_up_vote = ?", true).
		Where("grv.game_review_vote_game_review_id = ?", obj.ID).
		Find(&users).
		Error; err != nil {
			return nil, err
	}

	return users, nil
}

func (r *gameReviewResolver) DownVoters(ctx context.Context, obj *models.GameReview) ([]*models.User, error) {
	users := make([]*models.User, 0)

	if err := facades.UseDB().
		Joins("join game_review_votes grv on users.id = grv.game_review_vote_user_id").
		Where("is_up_vote = ?", false).
		Where("grv.game_review_vote_game_review_id = ?", obj.ID).
		Find(&users).
		Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *mutationResolver) CreateReview(ctx context.Context, gameID int64, content string, isRecommended bool) (*models.GameReview, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	game := new(models.Game)
	if err := facades.UseDB().First(game, gameID).Error; err != nil {
		return nil, err
	}

	review := &models.GameReview{
		Content:        content,
		IsRecommended:  isRecommended,
		GameReviewGame: *game,
		GameReviewUser: *user,
	}
	return review, facades.UseDB().Create(review).Error
}

func (r *mutationResolver) DeleteReview(ctx context.Context, id int64) (*models.GameReview, error) {
	review := new(models.GameReview)
	if err := facades.UseDB().First(review, id).Error; err != nil {
		return nil, err
	}
	return review, facades.UseDB().Delete(review).Error
}

func (r *mutationResolver) UpVoteReview(ctx context.Context, id int64) (*models.GameReview, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	review := new(models.GameReview)
	if err := facades.UseDB().Debug().First(review, id).Error; err != nil {
		return nil, err
	}
	log.Print(review)
	vote := new(models.GameReviewVote)
	if err := facades.UseDB().Debug().
		Where("game_review_vote_user_id = ?", user.ID).
		Where("game_review_vote_game_review_id = ?", review.ID).
		First(vote).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return review, facades.UseDB().Debug().Create(&models.GameReviewVote{
			GameReviewVoteGameReview: *review,
			GameReviewVoteUser:       *user,
			IsUpVote:                 true,
		}).Error
	}

	if vote.IsUpVote {
		return review, facades.UseDB().Debug().Delete(vote).Error
	}

	vote.IsUpVote = !vote.IsUpVote
	return review, facades.UseDB().Debug().Save(vote).Error
}

func (r *mutationResolver) DownVoteReview(ctx context.Context, id int64) (*models.GameReview, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	review := new(models.GameReview)
	if err := facades.UseDB().Debug().Debug().First(review, id).Error; err != nil {
		return nil, err
	}

	vote := new(models.GameReviewVote)
	if err := facades.UseDB().Debug().
		Where("game_review_vote_user_id = ?", user.ID).
		Where("game_review_vote_game_review_id = ?", review.ID).
		First(vote).
		Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		return review, facades.UseDB().Debug().Create(&models.GameReviewVote{
			GameReviewVoteGameReview: *review,
			GameReviewVoteUser:       *user,
			IsUpVote:                 false,
		}).Error
	}

	if !vote.IsUpVote {
		return review, facades.UseDB().Debug().Delete(vote).Error
	}

	vote.IsUpVote = !vote.IsUpVote
	return review, facades.UseDB().Debug().Save(vote).Error
}

// GameReview returns generated.GameReviewResolver implementation.
func (r *Resolver) GameReview() generated.GameReviewResolver { return &gameReviewResolver{r} }

type gameReviewResolver struct{ *Resolver }
