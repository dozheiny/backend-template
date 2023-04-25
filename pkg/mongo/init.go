package mongo

import (
	"context"
	"github.com/dozheiny/backend-template/config"
	"github.com/dozheiny/backend-template/database"
	"go.mongodb.org/mongo-driver/bson"
)

// Init will make initialize mongoDB client for future usages.
func Init(ctx context.Context) *Base {

	base := new(Base)
	base.Client = database.Client
	base.Context = ctx

	return base
}

// SetDatabase given a DBName and connect to database
// if there's no database or DBName is wrong, it returns error.
func (b *Base) SetDatabase() (*Base, error) {

	dbName, err := config.Get("DB_NAME")
	if err != nil {
		return nil, err
	}

	names, err := b.Client.ListDatabaseNames(b.Context, bson.D{})
	if err != nil {
		return nil, err
	}

	dbNames := make(map[string]bool)
	for number := range names {
		dbNames[names[number]] = true
	}

	if dbNames[dbName] {
		b.Database = b.Client.Database(dbName)
		b.Client = b.Database.Client()
		b.DataBaseName = dbName
		return b, nil
	}

	return nil, dbNameNotFound
}

// SetCollectionName sets collection name for future requests,
// if collection name doesn't exist, it creates collection name.
func (b *Base) SetCollectionName(collectionName string) (*Base, error) {
	names, err := b.Client.Database(b.DataBaseName).ListCollectionNames(b.Context, bson.D{})
	if err != nil {
		return nil, err
	}

	collectionNames := make(map[string]bool)
	for number := range names {
		collectionNames[names[number]] = true
	}

	if collectionNames[collectionName] {
		b.Collection = b.Client.Database(b.DataBaseName).Collection(collectionName)
		b.Database = b.Collection.Database()
		b.Client = b.Database.Client()
		return b, nil
	}

	if err := b.Database.CreateCollection(b.Context, collectionName); err != nil {
		return nil, err
	}

	b.Collection = b.Client.Database(b.DataBaseName).Collection(collectionName)
	b.Database = b.Collection.Database()
	b.Client = b.Database.Client()
	return b, nil
}
