package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myapp/graph/generated"
	"myapp/graph/model"
	"myapp/service"
)

func (r *aboutOpsResolver) CreateEducation(ctx context.Context, obj *model.AboutOps, input model.NewEducation) (*model.Education, error) {
	return service.CreateEducation(input), nil
}

func (r *aboutOpsResolver) CreateWorkingExperience(ctx context.Context, obj *model.AboutOps, input model.NewWorkingExperience) (*model.WorkingExperience, error) {
	return service.CreateWorkingExperience(input), nil
}

func (r *aboutOpsResolver) CreateBuiltProjects(ctx context.Context, obj *model.AboutOps, input model.NewBuiltProject) (*model.BuiltProject, error) {
	return service.CreateBuiltProject(input), nil
}

func (r *aboutOpsResolver) UpdateEducation(ctx context.Context, obj *model.AboutOps, input model.EditedEducation) (*model.Education, error) {
	return service.UpdateEducationByID(input)
}

func (r *aboutOpsResolver) UpdateWorkingExperience(ctx context.Context, obj *model.AboutOps, input model.EditedWorkingExperience) (*model.WorkingExperience, error) {
	return service.UpdateWorkingExperienceByID(input)
}

func (r *aboutOpsResolver) UpdateBuiltProjects(ctx context.Context, obj *model.AboutOps, input model.EditedBuiltProject) (*model.BuiltProject, error) {
	return service.UpdateBuiltProjectByID(input)
}

func (r *aboutOpsResolver) DeleteEducation(ctx context.Context, obj *model.AboutOps, id int) (int, error) {
	return service.DeleteEducationByID(id)
}

func (r *aboutOpsResolver) DeleteWorkingExprerience(ctx context.Context, obj *model.AboutOps, id int) (int, error) {
	return service.DeleteWorkingExperienceByID(id)
}

func (r *aboutOpsResolver) DeleteBuiltProjects(ctx context.Context, obj *model.AboutOps, id int) (int, error) {
	return service.DeleteBuiltProjectByID(id)
}

func (r *aboutOpsResolver) CreateEducationCategory(ctx context.Context, obj *model.AboutOps, input model.NewEducationCategory) (*model.EducationCategory, error) {
	return service.CreateEducationCategory(input),nil
}

func (r *aboutOpsResolver) UpdateEducationCategory(ctx context.Context, obj *model.AboutOps, input model.EditedEducationCategory) (*model.EducationCategory, error) {
	return service.UpdateEducationCategoryByID(input)
}

func (r *aboutOpsResolver) DeleteEducationCategory(ctx context.Context, obj *model.AboutOps, id int) (int, error) {
	return service.DeleteEducationCategoryByID(id)
}

func (r *aboutQueryResolver) Education(ctx context.Context, obj *model.AboutQuery, status []*int) ([]*model.EducationCategory, error) {
	return service.GetEducationCategory(status), nil
}

func (r *aboutQueryResolver) WorkingExperience(ctx context.Context, obj *model.AboutQuery, status []*int) ([]*model.WorkingExperience, error) {
	return service.GetWorkingExperience(status), nil
}

func (r *aboutQueryResolver) BuiltProjects(ctx context.Context, obj *model.AboutQuery, status []*int) ([]*model.BuiltProject, error) {
	return service.GetBuiltProjects(status), nil
}

func (r *aboutQueryResolver) EducationWithoutCategory(ctx context.Context, obj *model.AboutQuery) ([]*model.Education, error) {
	return service.GetAllEducation(), nil
}

// AboutOps returns generated.AboutOpsResolver implementation.
func (r *Resolver) AboutOps() generated.AboutOpsResolver { return &aboutOpsResolver{r} }

// AboutQuery returns generated.AboutQueryResolver implementation.
func (r *Resolver) AboutQuery() generated.AboutQueryResolver { return &aboutQueryResolver{r} }

type aboutOpsResolver struct{ *Resolver }
type aboutQueryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *aboutOpsResolver) Create(ctx context.Context, obj *model.AboutOps) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
