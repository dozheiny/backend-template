package mongo_test

import (
	"context"
	"github.com/dozheiny/backend-template/database"
	"github.com/dozheiny/backend-template/pkg/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

type model struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
	Code int                `bson:"code"`
}

func (m *model) GetCollectionName() string {
	return "this is test"
}

func (m *model) GetID() interface{} {
	return m.ID
}

func (m *model) SetID(id interface{}) {
	m.ID = id.(primitive.ObjectID)
}

func TestBase_Insert(t *testing.T) {

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

	id, err := base.Insert(record)
	if err != nil {
		t.Errorf("Gotcha Error on insert Record %s", err.Error())
	}

	output := new(model)

	err = base.FindOne(output, bson.D{{Key: "name", Value: bson.D{{Key: "$regex", Value: "hello"}}}})
	if err != nil {
		t.Errorf("got Error on find record: %s", err.Error())
	}

	if record.Name != output.Name {
		t.Errorf("returned object ID from insert and find is not the same. except: %s, actual: %s",
			id.String(), output.ID.String())
	}

}
