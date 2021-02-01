package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
)

func (r *communityResolver) ImageAndVideo(ctx context.Context, obj *models.Community, id int64) (*models.CommunityImageAndVideo, error) {
	obj.CommunityImageAndVideo = models.CommunityImageAndVideo{}
	return &obj.CommunityImageAndVideo, facades.UseDB().First(&obj.CommunityImageAndVideo, id).Error
}

func (r *communityResolver) ImagesAndVideos(ctx context.Context, obj *models.Community) ([]*models.CommunityImageAndVideo, error) {
	obj.CommunityImagesAndVideos = make([]*models.CommunityImageAndVideo, 0)
	return obj.CommunityImagesAndVideos, facades.UseDB().
		Find(&obj.CommunityImagesAndVideos).
		Error
}

func (r *communityResolver) Review(ctx context.Context, obj *models.Community, id int64) (*models.GameReview, error) {
	obj.CommunityReview = models.GameReview{}
	return &obj.CommunityReview, facades.UseDB().First(&obj.CommunityReview, id).Error
}

func (r *communityResolver) Reviews(ctx context.Context, obj *models.Community) ([]*models.GameReview, error) {
	obj.CommunityReviews = make([]*models.GameReview, 0)
	return obj.CommunityReviews, facades.UseDB().
		Find(&obj.CommunityReviews).
		Error
}

func (r *communityResolver) Discussion(ctx context.Context, obj *models.Community, id int64) (*models.CommunityDiscussion, error) {
	obj.CommunityDiscussion = models.CommunityDiscussion{}
	return &obj.CommunityDiscussion, facades.UseDB().First(&obj.CommunityDiscussion, id).Error
}

func (r *communityResolver) Discussions(ctx context.Context, obj *models.Community) ([]*models.CommunityDiscussion, error) {
	obj.CommunityDiscussions = make([]*models.CommunityDiscussion, 0)
	return obj.CommunityDiscussions, nil
}

func (r *queryResolver) Community(ctx context.Context) (*models.Community, error) {
	return new(models.Community), nil
}

// Community returns generated.CommunityResolver implementation.
func (r *Resolver) Community() generated.CommunityResolver { return &communityResolver{r} }

type communityResolver struct{ *Resolver }
