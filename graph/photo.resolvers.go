package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/dataloaders"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/utils"
)

func (r *imagesCategoryResolver) Photos(ctx context.Context, obj *model.ImagesCategory, status []int) ([]*model.Photo, error) {
	return dataloaders.CtxLoaders(ctx).Photo.Load(&model.PhotoLoaderParams{
		ID:     obj.ID,
		Status: status,
	})
}

func (r *photoResolver) Category(ctx context.Context, obj *model.Photo) (*model.ImagesCategory, error) {
	return dataloaders.CtxLoaders(ctx).ImageCategory.Load(obj.ImagesCategoryID)
}

func (r *photoResolver) Sizes(ctx context.Context, obj *model.Photo) (*model.File, error) {
	return utils.Resize(obj.Src), nil
}

func (r *topImageResolver) Sizes(ctx context.Context, obj *model.TopImage) (*model.File, error) {
	return utils.Resize(obj.Src), nil
}

// ImagesCategory returns generated.ImagesCategoryResolver implementation.
func (r *Resolver) ImagesCategory() generated.ImagesCategoryResolver {
	return &imagesCategoryResolver{r}
}

// Photo returns generated.PhotoResolver implementation.
func (r *Resolver) Photo() generated.PhotoResolver { return &photoResolver{r} }

// TopImage returns generated.TopImageResolver implementation.
func (r *Resolver) TopImage() generated.TopImageResolver { return &topImageResolver{r} }

type imagesCategoryResolver struct{ *Resolver }
type photoResolver struct{ *Resolver }
type topImageResolver struct{ *Resolver }
