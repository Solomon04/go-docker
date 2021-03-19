package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/solomon04/go-docker/graph/generated"
	"github.com/solomon04/go-docker/graph/model"
)

func (r *mutationResolver) RegisterUser(ctx context.Context, registration *model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AuthenticateWithPassword(ctx context.Context, payload *model.BasicLogin) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AuthenticateWithGoogle(ctx context.Context, payload *model.GoogleOAuth) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AuthenticateWithApple(ctx context.Context, payload *model.AppleOAuth) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Logout(ctx context.Context) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userInfo *model.UpdateUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSettings(ctx context.Context, userSettingsInfo *model.UpdateSettings) (*model.UserSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) JoinLocation(ctx context.Context, slug string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LeaveLocation(ctx context.Context, slug string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CommentOnLocation(ctx context.Context, message string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) JoinEvent(ctx context.Context, eventID string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) LeaveEvent(ctx context.Context, eventID string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CommentOnEvent(ctx context.Context, message string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, uuid *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, name *string, uuid *string, email *string, id *int) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Location(ctx context.Context, slug *string) (*model.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Locations(ctx context.Context) ([]*model.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RecommendedLocations(ctx context.Context) ([]*model.Location, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Event(ctx context.Context, eventID *string) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RecommendedEvents(ctx context.Context) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
