package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *homeOpsResolver) Update(ctx context.Context, obj *model.HomeOps, input model.EditedHome) (*model.Home, error) {
	return service.UpdateHomeInfo(input)
}

func (r *homeQueryResolver) GetInfo(ctx context.Context, obj *model.HomeQuery) (*model.Home, error) {
	return service.GetHomeInfo(), nil
}

// HomeOps returns generated.HomeOpsResolver implementation.
func (r *Resolver) HomeOps() generated.HomeOpsResolver { return &homeOpsResolver{r} }

// HomeQuery returns generated.HomeQueryResolver implementation.
func (r *Resolver) HomeQuery() generated.HomeQueryResolver { return &homeQueryResolver{r} }

type homeOpsResolver struct{ *Resolver }
type homeQueryResolver struct{ *Resolver }
