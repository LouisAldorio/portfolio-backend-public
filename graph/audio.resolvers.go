package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
	"myapp/utils"
)

func (r *audioResolver) Sizes(ctx context.Context, obj *model.Audio) (*model.File, error) {
	return utils.Resize(obj.Cover), nil
}

func (r *audioOpsResolver) Create(ctx context.Context, obj *model.AudioOps, input model.NewAudio) (*model.Audio, error) {
	return service.CreateAudio(input), nil
}

func (r *audioOpsResolver) Update(ctx context.Context, obj *model.AudioOps, input model.EditedAudio) (*model.Audio, error) {
	return service.UpdateAudioByID(input)
}

func (r *audioOpsResolver) Delete(ctx context.Context, obj *model.AudioOps, id int) (int, error) {
	return service.DeleteAudioByID(id)
}

// Audio returns generated.AudioResolver implementation.
func (r *Resolver) Audio() generated.AudioResolver { return &audioResolver{r} }

// AudioOps returns generated.AudioOpsResolver implementation.
func (r *Resolver) AudioOps() generated.AudioOpsResolver { return &audioOpsResolver{r} }

type audioResolver struct{ *Resolver }
type audioOpsResolver struct{ *Resolver }
