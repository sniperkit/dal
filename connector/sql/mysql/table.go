package mysql

import (
	"errors"

	"github.com/sniperkit/databases"
	"github.com/sniperkit/databases/type/relational"
)

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	_, ok := db.(*mysqlDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a mysql database instance")
	}
	return relational.NewTable(db, tableName)
}
