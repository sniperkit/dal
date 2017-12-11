package bigquery

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
	"github.com/sniperkit/databases"
)

type bigqueryConfig struct {
	PemPath             string
	ServiceAccountEmail string
	ServiceAccountID    string
	Secret              string
	Dataset             string
}

// Config ...
var (
	Config = &bigqueryConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (bigqueryConfig) ConfigName() string {
	return "BigQuery"
}

// SetDefaults ...
func (a *bigqueryConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *bigqueryConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

// Wait ...
func (c bigqueryConfig) Wait() {
	<-c.done
}

// String ...
func (c bigqueryConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c bigqueryConfig) Debug() {
	log.Debug("BigQuery Config = ", c)
}

func init() {
	config.Register(Config)
}
