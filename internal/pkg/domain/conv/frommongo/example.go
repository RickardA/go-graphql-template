package frommongo

import (
	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
	mongo "github.com/RickardA/go-graphql-template/internal/pkg/repository/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ExampleID(input primitive.ObjectID) domain.ExampleID {
	return domain.ExampleID(input.Hex())
}

func Example(input mongo.OutputExample) (domain.Example, error) {
	return domain.Example{
		ID:        ExampleID(input.ID),
		Variable1: input.Variable1,
		Variable2: input.Variable2,
	}, nil
}
