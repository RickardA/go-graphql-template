package mongo

import (
	"context"

	"github.com/RickardA/go-graphql-template/internal/pkg/domain"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain/conv/frommongo"
	"github.com/RickardA/go-graphql-template/internal/pkg/domain/conv/intomongo"
	"github.com/RickardA/go-graphql-template/internal/pkg/repository"
	mongo "github.com/RickardA/go-graphql-template/internal/pkg/repository/mongo/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var _ repository.ExampleRepository = &Client{}

var ExampleCollectionName = "examples"

func (c *Client) GetByID(ctx context.Context, id domain.ExampleID) (domain.Example, error) {
	coll := c.db.Database("db").Collection(ExampleCollectionName)

	convID, err := intomongo.ExampleID(id)

	if err != nil {
		log.WithField("id", id).Error("Could not convert domain ID to mongo ID")
		return domain.Example{}, err
	}

	result := coll.FindOne(c.ctx, bson.M{"_id": convID})

	var res mongo.OutputExample

	bytes, err := result.DecodeBytes()
	if err != nil {
		log.WithField("id", id).Error("Could not decode result")
		return domain.Example{}, err
	}

	err = bson.Unmarshal(bytes, &res)

	if err != nil {
		log.WithField("id", id).Error("Could not unmarshal result")
		return domain.Example{}, err
	}

	return frommongo.Example(res)
}

func (c *Client) GetAll(ctx context.Context) ([]domain.Example, error) {
	// coll := c.db.Database("db").Collection(ExampleCollectionName)

	// result := coll.Find(c.ctx)

	// var res mongo.OutputExample

	// bytes, err := result.DecodeBytes()
	// if err != nil {
	// 	log.WithField("id", id).Error("Could not decode result")
	// 	return domain.Example{}, err
	// }

	// err = bson.Unmarshal(bytes, &res)

	// if err != nil {
	// 	log.WithField("id", id).Error("Could not unmarshal result")
	// 	return domain.Example{}, err
	// }

	// return frommongo.Example(res)
	return []domain.Example{}, nil
}

func (c *Client) CreateExample(ctx context.Context, input domain.Example) (domain.ExampleID, error) {
	ipt, err := intomongo.Example(input)

	if err != nil {
		return "", err
	}

	obj, err := bson.Marshal(ipt)
	if err != nil {
		return "", err
	}

	coll := c.db.Database("db").Collection(ExampleCollectionName)

	result, err := coll.InsertOne(c.ctx, obj)

	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		if err != nil {
			return "", err
		}

		return frommongo.ExampleID(oid), nil
	}

	return "", repository.ErrCouldNotGetObjectID
}

func (c *Client) UpdateExample(ctx context.Context, input domain.Example) (domain.Example, error) {
	coll := c.db.Database("db").Collection(ExampleCollectionName)

	convID, err := intomongo.ExampleID(input.ID)

	if err != nil {
		log.WithField("id", input.ID).Error("Could not convert domain ID to mongo ID")
		return domain.Example{}, err
	}

	mongoObj, err := intomongo.Example(input)

	if err != nil {
		log.WithField("id", input.ID).Error("Could not convert domain obj to mongo obj")
		return domain.Example{}, err
	}

	updateBytes, err := bson.Marshal(mongoObj)

	if err != nil {
		log.WithField("id", input.ID).Error("Could not convert input to json")
		return domain.Example{}, err
	}

	var updateJSON bson.M
	err = bson.Unmarshal(updateBytes, &updateJSON)

	if err != nil {
		log.WithField("id", input.ID).Error("Could not unmarshal")
		return domain.Example{}, err
	}

	result, err := coll.UpdateOne(c.ctx, bson.M{"_id": convID}, bson.M{"$set": updateJSON})

	if err != nil {
		log.WithField("id", input.ID).Error("Could not update in db")
		return domain.Example{}, err
	}

	if result.MatchedCount == 0 {
		log.WithField("id", input.ID).Error("Could not update in db, id not found")
		return domain.Example{}, repository.ErrIDNotFound
	}

	return c.GetByID(ctx, input.ID)
}

func (c *Client) DeleteExample(ctx context.Context, id domain.ExampleID) (domain.ExampleID, error) {
	coll := c.db.Database("db").Collection(ExampleCollectionName)

	convID, err := intomongo.ExampleID(id)

	if err != nil {
		return "", err
	}

	_, err = coll.DeleteOne(c.ctx, bson.M{"_id": convID})

	if err != nil {
		return "", err
	}

	return id, nil
}
