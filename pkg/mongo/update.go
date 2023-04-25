package mongo

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UpdateByID will update record with ID.
func (b *Base) UpdateByID(recorder Recorder, id *primitive.ObjectID, query interface{}) error {
	b, err := b.SetCollectionName(recorder.GetCollectionName())
	if err != nil {
		return err
	}

	if _, err := b.Collection.UpdateByID(b.Context, id, query); err != nil {
		return err
	}

	return nil
}

// UpdateByID is shortcut for Base.UpdateByID.
func UpdateByID(c *fiber.Ctx, recorder Recorder, id *primitive.ObjectID, query interface{}) error {
	base, err := Init(c.Context()).SetDatabase()
	if err != nil {
		return err
	}

	if err := base.UpdateByID(recorder, id, query); err != nil {
		return err
	}

	return nil
}
