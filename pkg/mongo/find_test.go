package mongo_test

import (
	"context"
	"github.com/dozheiny/backend-template/database"
	"github.com/dozheiny/backend-template/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestBase_FindOne(t *testing.T) {

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

	if _, err := base.Insert(record); err != nil {
		t.Errorf("Gotcha Error on insert Record %s", err.Error())
	}

	output := new(model)

	err = base.FindOne(output, bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: "hello"}}}})
	if err != nil {
		t.Errorf("got Error on find record: %s", err.Error())
	}

	if output.Name != "hello" {
		t.Errorf("result and exception is not compare, result: %s, except: %s", output.Name, "hello")
	}

}

func TestBase_FindAll(t *testing.T) {
	ctx := context.Background()

	if err := database.Connection(ctx); err != nil {
		t.Errorf("Gotcha Error on set connection %s", err.Error())
	}

	base := mongo.Init(ctx)
	base, err := base.SetDatabase()
	if err != nil {
		t.Errorf("Gotcha Error on set Database %s", err.Error())
	}

	records := make([]model, 0)

	record1 := new(model)
	record1.ID = primitive.NewObjectID()
	record1.Name = "hello"

	record2 := new(model)
	record2.ID = primitive.NewObjectID()
	record2.Name = "mello"

	records = append(records, *record1, *record2)

	for number := range records {
		if _, err := base.Insert(&records[number]); err != nil {
			t.Errorf("Gotcha Error on insert Record %s", err.Error())
		}
	}

	cursor, _, err := base.FindAll(&model{}, bson.D{}, nil, nil)
	if err != nil {
		t.Errorf("gotcha Error on find all: %s", err.Error())
	}

	for cursor.Next(ctx) {
		o := new(model)
		if err := cursor.Decode(o); err != nil {
			t.Errorf("gotcha error on find all %s", err.Error())
		}
	}

}
