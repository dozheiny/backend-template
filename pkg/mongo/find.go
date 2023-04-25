package mongo

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math"
)

// FindOne gives Recorder and query and do query job,
// It fetches records and return Recorder again.
func (b *Base) FindOne(recorder Recorder, query interface{},
	opts ...*options.FindOneOptions) error {

	b, err := b.SetCollectionName(recorder.GetCollectionName())
	if err != nil {
		return err
	}

	if len(opts) == 0 {
		if err := b.Collection.FindOne(b.Context, query).Decode(recorder); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return NoDocuments
			}

			return err
		}
	} else {
		if err := b.Collection.FindOne(b.Context, query, opts[0]).Decode(recorder); err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return NoDocuments
			}

			return err
		}
	}

	return nil
}

// FindAll will execute query and return everything contains query.
func (b *Base) FindAll(recorder Recorder, query interface{},
	page, perPage *int, opts ...*options.FindOptions) (*mongo.Cursor, *Counters, error) {

	b, err := b.SetCollectionName(recorder.GetCollectionName())
	if err != nil {
		return nil, nil, err
	}

	totalDocs, err := b.Collection.CountDocuments(b.Context, query)
	if err != nil {
		return nil, nil, err
	}

	var totalPages int
	p := new(Paginate)
	if perPage != nil && page != nil {
		p.Limit = int64(*perPage)
		p.Page = int64(*page)

		totalPages = int(math.Ceil(float64(totalDocs) / float64(*perPage)))
	}

	var cursor *mongo.Cursor

	if len(opts) > 0 {
		cursor, err = b.Collection.Find(b.Context, query, p.Paginate(), opts[0])
	} else {
		cursor, err = b.Collection.Find(b.Context, query, p.Paginate())
	}

	if totalDocs == 0 {
		return nil, nil, NoDocuments
	}

	return cursor, &Counters{
		TotalPage:     totalPages,
		TotalDocument: int(totalDocs),
	}, err
}

// FindAll is a shortcut for Base.FindAll.
func FindAll(c *fiber.Ctx, recorder Recorder,
	query interface{}, page, perPage *int, opts ...*options.FindOptions) (*mongo.Cursor, *Counters, error) {

	base, err := Init(c.Context()).SetDatabase()
	if err != nil {
		return nil, nil, err
	}

	if len(opts) > 0 {
		return base.FindAll(recorder, query, page, perPage, opts[0])
	}

	return base.FindAll(recorder, query, page, perPage, nil)
}

// FindOne is a shortcut for Base.FindOne.
func FindOne(c *fiber.Ctx, recorder Recorder, query interface{}, opts ...*options.FindOneOptions) error {

	base, err := Init(c.Context()).SetDatabase()
	if err != nil {
		return err
	}

	if len(opts) > 0 {
		return base.FindOne(recorder, query, opts[0])
	}

	return base.FindOne(recorder, query)
}
