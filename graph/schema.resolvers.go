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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
