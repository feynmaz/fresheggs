package mongo

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once   sync.Once
	client *mongo.Client
	err    error
)

func GetMongoClient(ctx context.Context, mongoURI string) (*mongo.Client, error) {
	once.Do(func() {
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
		if err != nil {
			return
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			client.Disconnect(ctx)
			client = nil
			return
		}
	})

	return client, err
}
