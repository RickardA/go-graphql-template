package intomongo

import (
	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
	mongo "github.com/RickardA/go-graphql-template/internal/pkg/repository/mongo/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ExampleID(input domain.ExampleID) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(string(input))
}

func Example(input domain.Example) (mongo.InputExample, error) {
	return mongo.InputExample{
		Variable1: input.Variable1,
		Variable2: input.Variable2,
	}, nil
}
