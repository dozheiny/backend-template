package mongo_test

import (
	"context"
	"github.com/dozheiny/backend-template/database"
	"github.com/dozheiny/backend-template/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"testing"
)

func TestInit(t *testing.T) {

	ctx := context.Background()

	if err := database.Connection(ctx); err != nil {
		t.Errorf("Gotcha Error on database connection: %s", err.Error())
	}

	base := mongo.Init(ctx)
	if base.Client != database.Client {
		t.Errorf("base client and database client are not the same.")
	}

}

func TestBase_SetDatabase(t *testing.T) {

	ctx := context.Background()

	if err := database.Connection(ctx); err != nil {
		t.Errorf("Gotcha Error on database connection: %s", err.Error())
	}

	base := mongo.Init(ctx)

	if _, err := base.SetDatabase(); err != nil {
		t.Errorf("Gotcha Error on create database. %s", err.Error())
	}

	dbNames := make(map[string]bool)

	names, err := database.Client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		t.Errorf("Gotcha Error on create database. %s", err.Error())
	}

	for numbers := range names {
		dbNames[names[numbers]] = true
	}

	dbName := os.Getenv("DB_NAME")

	if !dbNames[dbName] {
		t.Errorf("Database dosen't sets!")
	}

}

func TestBase_SetCollectionName(t *testing.T) {
	var err error
	var collectionName = "config"

	ctx := context.Background()

	if err := database.Connection(ctx); err != nil {
		t.Errorf("Gotcha Error on database connection: %s", err.Error())
	}

	base := mongo.Init(ctx)
	base, err = base.SetDatabase()
	if err != nil {
		t.Errorf("Gotcha Error on create database. %s", err.Error())
	}

	if _, err := base.SetCollectionName(collectionName); err != nil {
		t.Errorf("gotcha error on set collection name: %s", err.Error())
	}

	cNames := make(map[string]bool)
	names, err := database.Client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		t.Errorf("Gotcha Error on set collection name: %s", err.Error())
	}

	for number := range names {
		cNames[names[number]] = true
	}

	if !cNames[collectionName] {
		t.Errorf("collection name didn't set")
	}

}
