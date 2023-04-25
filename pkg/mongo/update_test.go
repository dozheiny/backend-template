package mongo_test

import (
	"context"
	"github.com/dozheiny/backend-template/database"
	"github.com/dozheiny/backend-template/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestUpdateByID(t *testing.T) {
	ctx := context.Background()

	if err := database.Connection(ctx); err != nil {
		t.Errorf("Gotcha Error on set connection %s", err.Error())
	}

	base := mongo.Init(ctx)
	base, err := base.SetDatabase()
	if err != nil {
		t.Errorf("Gotcha Error on set Database %s", err.Error())
	}

	record := new(model)
	record.ID = primitive.NewObjectID()
	record.Name = "hello"

	_, err = base.Insert(record)
	if err != nil {
		t.Errorf("Gotcha Error on insert Record %s", err.Error())
	}

	if err := base.UpdateByID(record, &record.ID,
		bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: "bello"}}}}); err != nil {
		t.Errorf("gotcha error on update record: %s", err.Error())
	}

}
