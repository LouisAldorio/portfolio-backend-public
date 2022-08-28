package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *galleryOpsResolver) CreateTopImage(ctx context.Context, obj *model.GalleryOps, input model.NewTopImage) (*model.TopImage, error) {
	return service.CreateTopImage(input), nil
}

func (r *galleryOpsResolver) CreateImageCategory(ctx context.Context, obj *model.GalleryOps, input model.NewImagesCategory) (*model.ImagesCategory, error) {
	return service.CreateImageCategory(input), nil
}

func (r *galleryOpsResolver) CreateImage(ctx context.Context, obj *model.GalleryOps, input model.NewPhoto) (*model.Photo, error) {
	return service.CreateImage(input), nil
}

func (r *galleryOpsResolver) UpdateTopImage(ctx context.Context, obj *model.GalleryOps, input model.EditedTopImage) (*model.TopImage, error) {
	return service.UpdateTopImageByID(input)
}

func (r *galleryOpsResolver) UpdateImageCategory(ctx context.Context, obj *model.GalleryOps, input model.EditedImagesCategory) (*model.ImagesCategory, error) {
	return service.UpdateImageCategoryByID(input)
}

func (r *galleryOpsResolver) UpdateImage(ctx context.Context, obj *model.GalleryOps, input model.EditedPhoto) (*model.Photo, error) {
	return service.UpdateImageByID(input)
}

func (r *galleryOpsResolver) DeleteTopImage(ctx context.Context, obj *model.GalleryOps, id int) (int, error) {
	return service.DeleteTopImagesByID(id)
}

func (r *galleryOpsResolver) DeleteImage(ctx context.Context, obj *model.GalleryOps, id int) (int, error) {
	return service.DeleteImageByID(id)
}

func (r *galleryOpsResolver) DeleteImageCategory(ctx context.Context, obj *model.GalleryOps, id int) (int, error) {
	return service.DeleteImageCategoryByID(id)
}

func (r *galleryOpsResolver) Video(ctx context.Context, obj *model.GalleryOps) (*model.VideoOps, error) {
	return &model.VideoOps{}, nil
}

func (r *galleryOpsResolver) Audio(ctx context.Context, obj *model.GalleryOps) (*model.AudioOps, error) {
	return &model.AudioOps{}, nil
}

func (r *galleryQueryResolver) CarouselPhotos(ctx context.Context, obj *model.GalleryQuery, status []*int) ([]*model.TopImage, error) {
	return service.GetCarouselTopImages(status), nil
}

func (r *galleryQueryResolver) TabsPhotos(ctx context.Context, obj *model.GalleryQuery, status []*int) ([]*model.ImagesCategory, error) {
	return service.GetImageCategory(status), nil
}

func (r *galleryQueryResolver) Videos(ctx context.Context, obj *model.GalleryQuery, status []*int) ([]*model.Video, error) {
	return service.GetVideos(status), nil
}

func (r *galleryQueryResolver) Audios(ctx context.Context, obj *model.GalleryQuery, status []*int) ([]*model.Audio, error) {
	return service.GetAudios(status), nil
}

func (r *galleryQueryResolver) PhotosWithoutCategory(ctx context.Context, obj *model.GalleryQuery) ([]*model.Photo, error) {
	return service.GetImageWithoutCategory(), nil
}

// GalleryOps returns generated.GalleryOpsResolver implementation.
func (r *Resolver) GalleryOps() generated.GalleryOpsResolver { return &galleryOpsResolver{r} }

// GalleryQuery returns generated.GalleryQueryResolver implementation.
func (r *Resolver) GalleryQuery() generated.GalleryQueryResolver { return &galleryQueryResolver{r} }

type galleryOpsResolver struct{ *Resolver }
type galleryQueryResolver struct{ *Resolver }
