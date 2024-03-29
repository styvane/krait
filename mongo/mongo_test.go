package mongo

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func TestNewClient(t *testing.T) {
	uri, ok := os.LookupEnv("MongoURI")
	if !ok {
		t.Skip("MongoURI not exported")
	}

	c := &config{uri}

	client, err := InitClient(c)
	if err != nil && client != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if client == nil && err == nil {
		t.Error("expected error found nil")
	}
}

func TestInsertOne(t *testing.T) {
	uri, ok := os.LookupEnv("MongoURI")
	if !ok {
		t.Skip("MongoURI not exported")
	}
	c := &config{uri}
	client, err := InitClient(c)
	assert.Nil(t, err)
	ctx := context.Background()

	name := "devtestxxxx"
	d := &DBEntity{name, name}

	err = client.InsertOne(ctx, d, bson.M{"test": "test"})
	assert.Nil(t, err)

	t.Cleanup(func() {
		ctx := context.Background()
		db := client.Database(name)
		db.Drop(ctx)
		client.Disconnect(ctx)
	})
}
