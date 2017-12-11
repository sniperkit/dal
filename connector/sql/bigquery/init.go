package bigquery

import (
	bigquery_client "github.com/dailyburn/bigquery/client"
)

type Database struct {
	*bigquery_client.Client
}

func New(config *Config) (*Databsse, error) {
	return &Database{client.New(
		config.PemPath,
		config.ServiceAccountEmail,
		config.ServiceAccountID,
		config.Secret,
	)}
}

func (d *Database) Mount(client *bigquery_client.Client) {
	return &Database{client}
}

/*
// New instantiates a new client with the given params and return a reference to it
func New(pemPath string, options ...func(*Client) error) *Client {
	c := Client{
		pemPath:        pemPath,
		RequestTimeout: defaultRequestTimeout,
	}

	c.PrintDebug = false

	for _, option := range options {
		err := option(&c)
		if err != nil {
			return nil
		}
	}

	return &c
}
*/
