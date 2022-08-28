package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *mutationResolver) Home(ctx context.Context) (*model.HomeOps, error) {
	return &model.HomeOps{}, nil
}

func (r *mutationResolver) About(ctx context.Context) (*model.AboutOps, error) {
	return &model.AboutOps{}, nil
}

func (r *mutationResolver) Profile(ctx context.Context) (*model.ProfileOps, error) {
	return &model.ProfileOps{}, nil
}

func (r *mutationResolver) Contact(ctx context.Context) (*model.ContactOps, error) {
	return &model.ContactOps{}, nil
}

func (r *mutationResolver) Gallery(ctx context.Context) (*model.GalleryOps, error) {
	return &model.GalleryOps{}, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.OwnerCredential) (*model.Owner, error) {
	return service.Login(input)
}

func (r *queryResolver) About(ctx context.Context) (*model.AboutQuery, error) {
	return &model.AboutQuery{}, nil
}

func (r *queryResolver) Profile(ctx context.Context) (*model.ProfileQuery, error) {
	return &model.ProfileQuery{}, nil
}

func (r *queryResolver) Contact(ctx context.Context) (*model.ContactQuery, error) {
	return &model.ContactQuery{}, nil
}

func (r *queryResolver) Gallery(ctx context.Context) (*model.GalleryQuery, error) {
	return &model.GalleryQuery{}, nil
}

func (r *queryResolver) Home(ctx context.Context) (*model.HomeQuery, error) {
	return &model.HomeQuery{}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
