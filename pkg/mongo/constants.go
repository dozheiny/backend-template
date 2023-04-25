package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Base struct {
	Client       *mongo.Client
	Database     *mongo.Database
	Collection   *mongo.Collection
	Context      context.Context
	DataBaseName string
}

type Paginate struct {
	Limit int64
	Page  int64
}

type Counters struct {
	TotalPage     int
	TotalDocument int
}

type Recorder interface {
	GetCollectionName() string
	GetID() interface{}
}

var (
	dbNameNotFound        = errors.New("database name not found, maybe db name is wrong")
	cannotConvertInsertID = errors.New("cannot convert insertID to object ID")
	cannotDeleteRecord    = errors.New("cannot delete record")
	NoDocuments           = mongo.ErrNoDocuments
)
