package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *skillOpsResolver) Create(ctx context.Context, obj *model.SkillOps, input model.NewSkill) (*model.Skill, error) {
	return service.CreateSkill(input), nil
}

func (r *skillOpsResolver) Update(ctx context.Context, obj *model.SkillOps, input model.EditedSkill) (*model.Skill, error) {
	return service.UpdateSkillByID(input)
}

func (r *skillOpsResolver) Delete(ctx context.Context, obj *model.SkillOps, id int) (int, error) {
	return service.DeleteSkillByID(id)
}

// SkillOps returns generated.SkillOpsResolver implementation.
func (r *Resolver) SkillOps() generated.SkillOpsResolver { return &skillOpsResolver{r} }

type skillOpsResolver struct{ *Resolver }
