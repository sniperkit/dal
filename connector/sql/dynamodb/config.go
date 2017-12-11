package dynamodb

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
	"github.com/sniperkit/databases"
)

type dynamodbConfig struct {
	Provider       string        `json:"provider" config:"database.provider"`
	Endpoints      []string      `json:"endpoints" config:"database.endpoints"`
	MaxConnections int           `json:"max_connections" config:"database.max_connections" default:"0"`
	done           chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	Config = &dynamodbConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (dynamodbConfig) ConfigName() string {
	return "DynamoDB"
}

// SetDefaults ...
func (a *dynamodbConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *dynamodbConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

// Wait ...
func (c dynamodbConfig) Wait() {
	<-c.done
}

// String ...
func (c dynamodbConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c dynamodbConfig) Debug() {
	log.Debug("DynamoDB Config = ", c)
}

func init() {
	config.Register(Config)
}
