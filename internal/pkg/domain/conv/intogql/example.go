package intogql

import (
	"github.com/RickardA/go-graphql-template/graph/model"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
)

func Example(input domain.Example) *model.GQExample {
	return &model.GQExample{
		ID:        string(input.ID),
		Variable1: input.Variable1,
		Variable2: input.Variable2,
	}
}
