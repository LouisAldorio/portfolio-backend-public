package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *videoOpsResolver) Create(ctx context.Context, obj *model.VideoOps, input model.NewVideo) (*model.Video, error) {
	return service.CreateVideo(input), nil
}

func (r *videoOpsResolver) Update(ctx context.Context, obj *model.VideoOps, input model.EditedVideo) (*model.Video, error) {
	return service.UpdateVideoByID(input)
}

func (r *videoOpsResolver) Delete(ctx context.Context, obj *model.VideoOps, id int) (int, error) {
	return service.DeleteVideoByID(id)
}

// VideoOps returns generated.VideoOpsResolver implementation.
func (r *Resolver) VideoOps() generated.VideoOpsResolver { return &videoOpsResolver{r} }

type videoOpsResolver struct{ *Resolver }
