package app

import (
	"context"

	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
)

func NewClient(repo Repository) Client {
	return Client{
		repository: repo,
		//Subs:       make(map[string]chan *model.GQConflict),
	}
}

type Client struct {
	repository Repository
	//Subs       map[string]chan *model.GQConflict
}

type Repository interface {
	ExampleRepository
}

type ExampleRepository interface {
	GetByID(ctx context.Context, id domain.ExampleID) (domain.Example, error)
	GetAll(ctx context.Context) ([]domain.Example, error)
	CreateExample(ctx context.Context, input domain.Example) (domain.ExampleID, error)
	UpdateExample(ctx context.Context, input domain.Example) (domain.Example, error)
	DeleteExample(ctx context.Context, id domain.ExampleID) (domain.ExampleID, error)
}
