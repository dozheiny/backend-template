package mongo

import "go.mongodb.org/mongo-driver/mongo/options"

// Paginate will paginate all records coming from DB
func (p *Paginate) Paginate() *options.FindOptions {
	skip := p.Page*p.Limit - p.Limit
	return &options.FindOptions{Limit: &p.Limit, Skip: &skip}
}
