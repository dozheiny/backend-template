// Package database is package for make connections to mongo pools.
package database

import (
	"context"
	"github.com/dozheiny/backend-template/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection will make connection to MongoDB database.
// if Connection can't connect or there's a problem, it returns error.
func Connection(ctx context.Context) error {
	uri, err := config.Get("MONGODB")
	if err != nil {
		return err
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	Client = client

	return nil
}
