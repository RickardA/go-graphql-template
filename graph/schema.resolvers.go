package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/RickardA/go-graphql-template/graph/generated"
	"github.com/RickardA/go-graphql-template/graph/model"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain/conv/fromgql"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain/conv/intogql"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func (r *mutationResolver) CreateExample(ctx context.Context, input model.GQCreateExample) (*string, error) {
	id, err := r.Client.CreateExample(ctx, fromgql.CreateExample(input))

	if err != nil {
		return nil, err
	}

	idAsString := string(id)

	return &idAsString, nil
}

func (r *mutationResolver) UpdateExample(ctx context.Context, input model.GQUpdateExample) (*model.GQExample, error) {
	res, err := r.Client.UpdateExample(ctx, fromgql.UpdateExample(input))

	if err != nil {
		return nil, err
	}

	return intogql.Example(res), err
}

func (r *mutationResolver) DeleteExample(ctx context.Context, id string) (*string, error) {
	res, err := r.Client.DeleteExample(ctx, domain.ExampleID(id))

	if err != nil {
		return nil, err
	}

	idAsString := string(res)

	return &idAsString, nil
}

func (r *queryResolver) GetByID(ctx context.Context, id string) (*model.GQExample, error) {
	res, err := r.Client.GetByID(ctx, domain.ExampleID(id))

	if err != nil {
		return nil, err
	}

	return intogql.Example(res), nil
}

func (r *queryResolver) GetAll(ctx context.Context) ([]*model.GQExample, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *subscriptionResolver) ExampleSubscription(ctx context.Context) (<-chan *model.GQExample, error) {

	subID := uuid.New()

	ch := make(chan *model.GQExample, 1)

	// Save subscription channel
	r.Client.Subs[subID.String()] = ch

	log.Info("Subscription recieved and handled")

	return ch, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
