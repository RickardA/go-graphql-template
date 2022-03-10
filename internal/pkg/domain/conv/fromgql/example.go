package fromgql

import (
	"github.com/RickardA/go-graphql-template/graph/model"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
)

func CreateExample(input model.GQCreateExample) domain.Example {
	return domain.Example{
		Variable1: input.Variable1,
		Variable2: input.Variable2,
	}
}

func UpdateExample(input model.GQUpdateExample) domain.Example {
	return domain.Example{
		ID:        domain.ExampleID(input.ID),
		Variable1: input.Variable1,
		Variable2: input.Variable2,
	}
}
