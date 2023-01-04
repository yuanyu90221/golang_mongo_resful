package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var MONGO_DATABASE string
var MONGO_COLLECTION string

func SetUpConnect(MONGO_URI string, ctx context.Context) {
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	Client, _ = mongo.Connect(ctx, clientOptions)
}
