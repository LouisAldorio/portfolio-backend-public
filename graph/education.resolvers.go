package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/dataloaders"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *educationResolver) Category(ctx context.Context, obj *model.Education) (*model.EducationCategory, error) {
	return service.GetEducationCategoryByID(obj.EducationCategoryID), nil
}

func (r *educationCategoryResolver) Nodes(ctx context.Context, obj *model.EducationCategory, status []int) ([]*model.Education, error) {
	return dataloaders.CtxLoaders(ctx).Education.Load(&model.EducationLoaderParams{
		ID:     obj.ID,
		Status: status,
	})
}

// Education returns generated.EducationResolver implementation.
func (r *Resolver) Education() generated.EducationResolver { return &educationResolver{r} }

// EducationCategory returns generated.EducationCategoryResolver implementation.
func (r *Resolver) EducationCategory() generated.EducationCategoryResolver {
	return &educationCategoryResolver{r}
}

type educationResolver struct{ *Resolver }
type educationCategoryResolver struct{ *Resolver }
