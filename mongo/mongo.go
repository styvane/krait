// package mongo implements all mongodb operations
package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

// MongoOperation wraps all database operations
type MongoOperation interface {
	InsertOne(ctx context.Context, dbname string, collection string, data interface{}) error
}

// MongoClient represent mongodb client
type MongoClient struct {
	log *zap.SugaredLogger
	*mongo.Client
}

// InitiClient create MongoClient with a connected mongo client.
// It also returns a error when connection to mongod fails or
// a logger configuration fails.
func InitClient(cfg *config) (*MongoClient, error) {
	opts := options.Client()
	client, err := mongo.NewClient(opts.ApplyURI(cfg.URI))
	if err != nil {
		return nil, err
	}

	log, _ := zap.NewProduction()
	sugar := log.Sugar()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	mc := &MongoClient{sugar, client}
	return mc, err
}

func (c *MongoClient) InsertOne(ctx context.Context, dbname string, collection string, data interface{}) error {
	db := c.Database(dbname)
	coll := db.Collection(collection)
	_, err := coll.InsertOne(ctx, data)
	return err
}
