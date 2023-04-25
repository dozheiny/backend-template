package mongo

import (
	"github.com/gofiber/fiber/v2"
)

// DeleteOne given delete record which query contains.
func (b *Base) DeleteOne(recorder Recorder, query interface{}) error {

	b, err := b.SetCollectionName(recorder.GetCollectionName())
	if err != nil {
		return err
	}

	res, err := b.Collection.DeleteOne(b.Context, query)
	if err != nil {
		return err
	}

	if res.DeletedCount < 1 {
		return cannotDeleteRecord
	}

	return nil
}

// DeleteOne is shortcut for Base.DeleteOne.
func DeleteOne(c *fiber.Ctx, recorder Recorder, query interface{}) error {
	base, err := Init(c.Context()).SetDatabase()
	if err != nil {
		return err
	}

	return base.DeleteOne(recorder, query)
}
