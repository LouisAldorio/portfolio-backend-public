package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *profileOpsResolver) ProgrammingLanguage(ctx context.Context, obj *model.ProfileOps) (*model.ProgrammingLanguageOps, error) {
	return &model.ProgrammingLanguageOps{}, nil
}

func (r *profileOpsResolver) Skill(ctx context.Context, obj *model.ProfileOps) (*model.SkillOps, error) {
	return &model.SkillOps{}, nil
}

func (r *profileQueryResolver) ProgrammingLanguages(ctx context.Context, obj *model.ProfileQuery, status []*int) ([]*model.ProgrammingLanguage, error) {
	return service.GetAllProgrammingLanguages(status), nil
}

func (r *profileQueryResolver) Skills(ctx context.Context, obj *model.ProfileQuery, status []*int) ([]*model.Skill, error) {
	return service.GetAllSkills(status), nil
}

// ProfileOps returns generated.ProfileOpsResolver implementation.
func (r *Resolver) ProfileOps() generated.ProfileOpsResolver { return &profileOpsResolver{r} }

// ProfileQuery returns generated.ProfileQueryResolver implementation.
func (r *Resolver) ProfileQuery() generated.ProfileQueryResolver { return &profileQueryResolver{r} }

type profileOpsResolver struct{ *Resolver }
type profileQueryResolver struct{ *Resolver }
