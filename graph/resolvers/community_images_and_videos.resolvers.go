package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"math"

	"github.com/brandon-julio-t/tpa-web-backend/facades"
	"github.com/brandon-julio-t/tpa-web-backend/graph/generated"
	"github.com/brandon-julio-t/tpa-web-backend/graph/models"
	"github.com/brandon-julio-t/tpa-web-backend/middlewares"
)

func (r *communityImageAndVideoResolver) Comments(ctx context.Context, obj *models.CommunityImageAndVideo, page int64) (*models.CommunityImageAndVideoCommentPagination, error) {
	comments := make([]*models.CommunityImageAndVideoComment, 0)
	perPage := 10
	count := new(int64)

	if err := facades.UseDB().
		Model(new(models.CommunityImageAndVideoComment)).
		Where("community_image_and_video_id = ?", obj.ID).
		Count(count).
		Scopes(facades.UsePagination(int(page), perPage)).
		Find(&comments).
		Error; err != nil {
		return nil, err
	}

	return &models.CommunityImageAndVideoCommentPagination{
		Data:       comments,
		TotalPages: int64(math.Ceil(float64(*count) / float64(perPage))),
	}, nil
}

func (r *communityImageAndVideoResolver) Dislikes(ctx context.Context, obj *models.CommunityImageAndVideo) (int64, error) {
	count := new(int64)
	return *count, facades.UseDB().
		Model(new(models.CommunityImageAndVideoRating)).
		Where("community_image_and_video_id = ?", obj.ID).
		Where("is_like = ?", true).
		Count(count).
		Error
}

func (r *communityImageAndVideoResolver) IsDisliked(ctx context.Context, obj *models.CommunityImageAndVideo) (bool, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return false, err
	}

	count := new(int64)
	return *count > 0, facades.UseDB().
		Model(new(models.CommunityImageAndVideoRating)).
		Where("user_id = ?", user.ID).
		Where("community_image_and_video_id = ?", obj.ID).
		Where("is_like = ?", false).
		Count(count).
		Error
}

func (r *communityImageAndVideoResolver) IsLiked(ctx context.Context, obj *models.CommunityImageAndVideo) (bool, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return false, err
	}

	count := new(int64)
	return *count > 0, facades.UseDB().
		Model(new(models.CommunityImageAndVideoRating)).
		Where("user_id = ?", user.ID).
		Where("community_image_and_video_id = ?", obj.ID).
		Where("is_like = ?", true).
		Count(count).
		Error
}

func (r *communityImageAndVideoResolver) Likes(ctx context.Context, obj *models.CommunityImageAndVideo) (int64, error) {
	count := new(int64)
	return *count, facades.UseDB().
		Model(new(models.CommunityImageAndVideoRating)).
		Where("community_image_and_video_id = ?", obj.ID).
		Where("is_like = ?", false).
		Count(count).
		Error
}

func (r *communityImageAndVideoCommentResolver) CommunityImagesAndVideos(ctx context.Context, obj *models.CommunityImageAndVideoComment) (*models.CommunityImageAndVideo, error) {
	obj.CommunityImageAndVideo_ = models.CommunityImageAndVideo{}
	return &obj.CommunityImageAndVideo_, facades.UseDB().First(&obj.CommunityImageAndVideo_, obj.CommunityImageAndVideoID).Error
}

func (r *mutationResolver) CreateCommunityImagesAndVideos(ctx context.Context, input models.CreateCommunityImageAndVideo) (*models.CommunityImageAndVideo, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(input.File.File)
	if err != nil {
		return nil, err
	}

	entity := &models.CommunityImageAndVideo{
		Description: input.Description,
		Name:        input.Name,
		File_: models.AssetFile{
			File:        data,
			ContentType: input.File.ContentType,
		},
		User_: *user,
	}

	return entity, facades.UseDB().Debug().Create(entity).Error
}

func (r *mutationResolver) LikeCreateCommunityImagesAndVideos(ctx context.Context, imageAndVideoID int64) (*models.CommunityImageAndVideo, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	imageAndVideo := new(models.CommunityImageAndVideo)
	if err := facades.UseDB().First(imageAndVideo, imageAndVideoID).Error; err != nil {
		return nil, err
	}

	rating := new(models.CommunityImageAndVideoRating)
	if err := facades.UseDB().
		Where("user_id = ?", user.ID).
		Where("community_image_and_video_id = ?", imageAndVideo.ID).
		First(rating).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return imageAndVideo, facades.UseDB().Create(&models.CommunityImageAndVideoRating{
				IsLike:                  true,
				CommunityImageAndVideo_: *imageAndVideo,
				User_:                   *user,
			}).Error
		}

		return nil, err
	}

	if !rating.IsLike {
		rating.IsLike = !rating.IsLike
		return imageAndVideo, facades.UseDB().Save(rating).Error
	}

	return imageAndVideo, facades.UseDB().Delete(rating).Error
}

func (r *mutationResolver) DislikeCreateCommunityImagesAndVideos(ctx context.Context, imageAndVideoID int64) (*models.CommunityImageAndVideo, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	imageAndVideo := new(models.CommunityImageAndVideo)
	if err := facades.UseDB().First(imageAndVideo, imageAndVideoID).Error; err != nil {
		return nil, err
	}

	rating := new(models.CommunityImageAndVideoRating)
	if err := facades.UseDB().
		Where("user_id = ?", user.ID).
		Where("community_image_and_video_id = ?", imageAndVideo.ID).
		First(rating).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return imageAndVideo, facades.UseDB().Create(&models.CommunityImageAndVideoRating{
				IsLike:                  false,
				CommunityImageAndVideo_: *imageAndVideo,
				User_:                   *user,
			}).Error
		}

		return nil, err
	}

	if rating.IsLike {
		rating.IsLike = !rating.IsLike
		return imageAndVideo, facades.UseDB().Save(rating).Error
	}

	return imageAndVideo, facades.UseDB().Delete(rating).Error
}

func (r *mutationResolver) PostCommunityImagesAndVideosComment(ctx context.Context, imageAndVideoID int64, body string) (*models.CommunityImageAndVideoComment, error) {
	user, err := middlewares.UseAuth(ctx)
	if err != nil {
		return nil, err
	}

	imageAndVideo := new(models.CommunityImageAndVideo)
	if err := facades.UseDB().First(imageAndVideo, imageAndVideoID).Error; err != nil {
		return nil, err
	}

	entity := &models.CommunityImageAndVideoComment{
		Body:                    body,
		User_:                   *user,
		CommunityImageAndVideo_: *imageAndVideo,
	}

	return entity, facades.UseDB().Create(entity).Error
}

// CommunityImageAndVideo returns generated.CommunityImageAndVideoResolver implementation.
func (r *Resolver) CommunityImageAndVideo() generated.CommunityImageAndVideoResolver {
	return &communityImageAndVideoResolver{r}
}

// CommunityImageAndVideoComment returns generated.CommunityImageAndVideoCommentResolver implementation.
func (r *Resolver) CommunityImageAndVideoComment() generated.CommunityImageAndVideoCommentResolver {
	return &communityImageAndVideoCommentResolver{r}
}

type communityImageAndVideoResolver struct{ *Resolver }
type communityImageAndVideoCommentResolver struct{ *Resolver }
