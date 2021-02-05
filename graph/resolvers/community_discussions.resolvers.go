package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
)

func (r *communityDiscussionResolver) Comments(ctx context.Context, obj *models.CommunityDiscussion, page int64) (*models.CommunityDiscussionCommentPagination, error) {
	entities := make([]*models.CommunityDiscussionComment, 0)
	count := new(int64)
	perPage := 10

	if err := facades.UseDB().
		Model(new(models.CommunityDiscussionComment)).
		Where("community_discussion_id = ?", obj.ID).
		Count(count).
		Scopes(facades.UsePagination(int(page), perPage)).
		Find(&entities).
		Error; err != nil {
		return nil, err
	}

	return &models.CommunityDiscussionCommentPagination{
		Data:       entities,
		TotalPages: int64(math.Ceil(float64(*count) / float64(perPage))),
	}, nil
}

func (r *gameResolver) TopDiscussions(ctx context.Context, obj *models.Game) ([]*models.CommunityDiscussion, error) {
	discussions := make([]*models.CommunityDiscussion, 0)
	return discussions, facades.UseDB().
		Where("game_id = ?", obj.ID).
		Order("created_at desc").
		Limit(3).
		Find(&discussions).
		Error
}

func (r *gameResolver) Discussions(ctx context.Context, obj *models.Game) ([]*models.CommunityDiscussion, error) {
	discussions := make([]*models.CommunityDiscussion, 0)
	return discussions, facades.UseDB().Where("game_id = ?", obj.ID).Find(&discussions).Error
}

func (r *mutationResolver) PostCommunityDiscussion(ctx context.Context, input models.PostCommunityDiscussion) (*models.CommunityDiscussion, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	game := new(models.Game)
	if err := facades.UseDB().First(game, input.GameID).Error; err != nil {
		return nil, err
	}

	entity := &models.CommunityDiscussion{
		Body:  input.Body,
		Title: input.Title,
		Game_: *game,
		User_: *user,
	}
	return entity, facades.UseDB().Create(entity).Error
}

func (r *mutationResolver) PostCommunityDiscussionComment(ctx context.Context, input models.PostCommunityDiscussionComment) (*models.CommunityDiscussionComment, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	cd := new(models.CommunityDiscussion)
	if err := facades.UseDB().First(cd, input.CommunityDiscussionID).Error; err != nil {
		return nil, err
	}

	entity := &models.CommunityDiscussionComment{
		Body:                 input.Body,
		CommunityDiscussion_: *cd,
		User_:                *user,
	}
	return entity, facades.UseDB().Create(entity).Error
}

func (r *queryResolver) GameDiscussion(ctx context.Context, id int64) (*models.CommunityDiscussion, error) {
	entity := new(models.CommunityDiscussion)
	return entity, facades.UseDB().First(entity, id).Error
}

func (r *queryResolver) GameDiscussions(ctx context.Context, title string) ([]*models.Game, error) {
	games := make([]*models.Game, 0)
	return games, facades.UseDB().
		Where(
			"id in (?)",
			facades.UseDB().
				Model(new(models.Game)).
				Select("games.id").
				Joins("join community_discussions cd on games.id = cd.game_id").
				Group("games.id").
				Having("count(cd.id) > 0"),
		).
		Find(&games, "lower(title) like lower(?)", "%"+title+"%").
		Error
}

// CommunityDiscussion returns generated.CommunityDiscussionResolver implementation.
func (r *Resolver) CommunityDiscussion() generated.CommunityDiscussionResolver {
	return &communityDiscussionResolver{r}
}

type communityDiscussionResolver struct{ *Resolver }
