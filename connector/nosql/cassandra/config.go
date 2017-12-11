package cassandra

import (
	"github.com/k0kubun/pp"
	"github.com/sniperkit/config"
	"github.com/sniperkit/databases"
	"github.com/sniperkit/vipertags"
)

type cassandradbConfig struct {
	Provider        string        `json:"provider" config:"database.provider"`
	Endpoints       []string      `json:"endpoints" config:"database.endpoints"`
	Username        string        `json:"username" config:"database.username"`
	Password        string        `json:"password" config:"database.password"`
	Cert            string        `json:"cert" config:"database.cert"`
	InitialCapacity int           `json:"initial_capacity" config:"database.initial_capacity" default:"0"`
	MaxConnections  int           `json:"max_connections" config:"database.max_connections" default:"0"`
	done            chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	Config = &cassandradbConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (cassandradbConfig) ConfigName() string {
	return "MongoDB"
}

// SetDefaults ...
func (a *cassandradbConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *cassandradbConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = 2 * database.DefaultMaxConnections
	}
}

// Wait ...
func (c cassandradbConfig) Wait() {
	<-c.done
}

// String ...
func (c cassandradbConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c cassandradbConfig) Debug() {
	log.Debug("MongoDB Config = ", c)
}

func init() {
	config.Register(Config)
}
