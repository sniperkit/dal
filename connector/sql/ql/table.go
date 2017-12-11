package ql

import (
	"errors"

	"github.com/sniperkit/databases"
	"github.com/sniperkit/databases/type/relational"
)

// NewTable ...
func NewTable(db database.Database, tableName string) (database.Table, error) {
	_, ok := db.(*qlDatabase)
	if !ok {
		return nil, errors.New("invalid database input. Expecting a sqlite database instance")
	}
	return relational.NewTable(db, tableName)
}
