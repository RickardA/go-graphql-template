package app

import (
	"context"

	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain/conv/intogql"
)

var _ ExampleRepository = &Client{}

func (c *Client) GetByID(ctx context.Context, id domain.ExampleID) (domain.Example, error) {
	return c.repository.GetByID(ctx, id)
}

func (c *Client) GetAll(ctx context.Context) ([]domain.Example, error) {
	return c.repository.GetAll(ctx)
}

func (c *Client) CreateExample(ctx context.Context, input domain.Example) (domain.ExampleID, error) {

	createdID, err := c.repository.CreateExample(ctx, input)

	if err != nil {
		return domain.ExampleID(""), err
	}

	obj, err := c.GetByID(ctx, createdID)

	if err != nil {
		return domain.ExampleID(""), err
	}

	for _, sub := range c.Subs {
		sub <- intogql.Example(obj)
	}

	return c.repository.CreateExample(ctx, input)
}

func (c *Client) UpdateExample(ctx context.Context, input domain.Example) (domain.Example, error) {
	return c.repository.UpdateExample(ctx, input)
}

func (c *Client) DeleteExample(ctx context.Context, id domain.ExampleID) (domain.ExampleID, error) {
	return c.repository.DeleteExample(ctx, id)
}
