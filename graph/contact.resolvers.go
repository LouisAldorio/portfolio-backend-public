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

func (r *contactOpsResolver) Mail(ctx context.Context, obj *model.ContactOps, input model.MailInput) (string, error) {
	return service.CreateSubscriber(input), nil
}

func (r *contactOpsResolver) UpdateSocialMediaInfo(ctx context.Context, obj *model.ContactOps, input model.EditSocialMediInfo) (*model.SocialMediaInfo, error) {
	return service.UpdateSocialMediaInfo(input)
}

func (r *contactQueryResolver) SocialMediaInfo(ctx context.Context, obj *model.ContactQuery) (*model.SocialMediaInfo, error) {
	return service.GetSocialMediaInfo(), nil
}

func (r *contactQueryResolver) Messages(ctx context.Context, obj *model.ContactQuery) ([]*model.Mail, error) {
	return service.GetMessages(), nil
}

func (r *mailResolver) Subscriber(ctx context.Context, obj *model.Mail) (*model.Subscriber, error) {
	return dataloaders.CtxLoaders(ctx).Subscriber.Load(obj.SubscriberID)
}

// ContactOps returns generated.ContactOpsResolver implementation.
func (r *Resolver) ContactOps() generated.ContactOpsResolver { return &contactOpsResolver{r} }

// ContactQuery returns generated.ContactQueryResolver implementation.
func (r *Resolver) ContactQuery() generated.ContactQueryResolver { return &contactQueryResolver{r} }

// Mail returns generated.MailResolver implementation.
func (r *Resolver) Mail() generated.MailResolver { return &mailResolver{r} }

type contactOpsResolver struct{ *Resolver }
type contactQueryResolver struct{ *Resolver }
type mailResolver struct{ *Resolver }
