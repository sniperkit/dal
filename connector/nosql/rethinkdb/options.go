package rethinkdb

import (
	"context"

	"github.com/sniperkit/databases"
	// 	"github.com/sniperkit/databases/type/collection"
)

const (
	authKeyKey         = "github.com/sniperkit/databases/connector/nosql/rethinkdb/authKey"
	initialCapacityKey = "github.com/sniperkit/databases/connector/nosql/rethinkdb/initialCapacity"
)

// DefaultInitialCapacity ...
var (
	DefaultInitialCapacity = 10
)

// AuthKey ...
func AuthKey(s string) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, authKeyKey, s)
	}
}

// InitialCapacity ...
func InitialCapacity(n int) database.Option {
	return func(o *database.Options) {
		o.Context = context.WithValue(o.Context, initialCapacityKey, n)
	}
}
