package ql

import (
	"github.com/k0kubun/pp"
	"github.com/rai-project/config"
	"github.com/rai-project/vipertags"
	"github.com/sniperkit/databases"
)

type qldbConfig struct {
	Provider       string        `json:"provider" config:"database.provider"`
	Endpoints      []string      `json:"endpoints" config:"database.endpoints"`
	MaxConnections int           `json:"max_connections" config:"database.max_connections" default:"0"`
	done           chan struct{} `json:"-" config:"-"`
}

// Config ...
var (
	Config = &qldbConfig{
		done: make(chan struct{}),
	}
)

// ConfigName ...
func (qldbConfig) ConfigName() string {
	return "QL"
}

// SetDefaults ...
func (a *qldbConfig) SetDefaults() {
	vipertags.SetDefaults(a)
}

// Read ...
func (a *qldbConfig) Read() {
	defer close(a.done)
	vipertags.Fill(a)
	if a.MaxConnections == 0 {
		a.MaxConnections = database.DefaultMaxConnections
	}
}

// Wait ...
func (c qldbConfig) Wait() {
	<-c.done
}

// String ...
func (c qldbConfig) String() string {
	return pp.Sprintln(c)
}

// Debug ...
func (c qldbConfig) Debug() {
	log.Debug("QL Config = ", c)
}

func init() {
	config.Register(Config)
}
