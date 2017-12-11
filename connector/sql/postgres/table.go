package postgres

import (
	"errors"

	"github.com/sniperkit/databases"
	"github.com/sniperkit/databases/type/relational"
)

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	_, ok := db.(*postgresDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a postgres database instance")
	}
	return relational.NewTable(db, tableName)
}
