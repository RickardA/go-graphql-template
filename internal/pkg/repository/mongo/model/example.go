package mongo_model

import "go.mongodb.org/mongo-driver/bson/primitive"

type InputExample struct {
	Variable1 string
	Variable2 int
}

type OutputExample struct {
	ID        primitive.ObjectID `bson:"_id"`
	Variable1 string
	Variable2 int
}
