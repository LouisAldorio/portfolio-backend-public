package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *programmingLanguageOpsResolver) Create(ctx context.Context, obj *model.ProgrammingLanguageOps, input model.NewProgrammingLanguage) (*model.ProgrammingLanguage, error) {
	return service.CreateProgrammingLanguage(input), nil
}

func (r *programmingLanguageOpsResolver) Update(ctx context.Context, obj *model.ProgrammingLanguageOps, input model.EditedProgrammingLanguage) (*model.ProgrammingLanguage, error) {
	return service.UpdateProgrammingLanguageByID(input)
}

func (r *programmingLanguageOpsResolver) Delete(ctx context.Context, obj *model.ProgrammingLanguageOps, id int) (int, error) {
	return service.DeleteProgrammingLanguageByID(id)
}

// ProgrammingLanguageOps returns generated.ProgrammingLanguageOpsResolver implementation.
func (r *Resolver) ProgrammingLanguageOps() generated.ProgrammingLanguageOpsResolver {
	return &programmingLanguageOpsResolver{r}
}

type programmingLanguageOpsResolver struct{ *Resolver }
