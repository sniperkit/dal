package ql

import (
	"context"
	"path/filepath"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sniperkit/databases"
	"github.com/sniperkit/databases/type/relational"
	"upper.io/db.v3/ql"
)

const (
	gormDialect = "ql"
)

type qlDatabase struct {
	database.Database
}

// NewDatabase ...
func NewDatabase(databaseName string, opts ...database.Option) (database.Database, error) {

	options := database.Options{
		Endpoints:      Config.Endpoints,
		TLSConfig:      nil,
		MaxConnections: Config.MaxConnections,
		Context:        context.Background(),
	}

	for _, o := range opts {
		o(&options)
	}

	connectionURL := ql.ConnectionURL{
		Database: filepath.Join(options.Endpoints[0], databaseName),
	}

	d, err := relational.NewDatabase(gormDialect, databaseName, connectionURL, options)
	if err != nil {
		return nil, err
	}
	return &sqliteDatabase{d}, nil
}

// String ...
func (conn *qlDatabase) String() string {
	return "QL"
}
