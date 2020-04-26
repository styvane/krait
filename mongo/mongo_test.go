package mongo

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

func TestNewClient(t *testing.T) {

	c := &config{"mongodb://locahost:20017"}

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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbname := "devtestxxxx"
	err = client.InsertOne(ctx, dbname, "devtestxxxx", bson.M{"test": "test"})
	assert.Nil(t, err)
	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		db := client.Database(dbname)
		db.Drop(ctx)
		client.Disconnect(ctx)
	})
}
