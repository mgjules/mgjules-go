package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
)

type EdgeDB struct {
	client *edgedb.Client
}

func New(ctx context.Context, dsn string) (*EdgeDB, error) {
	client, err := edgedb.CreateClientDSN(ctx, dsn, edgedb.Options{})
	if err != nil {
		return nil, fmt.Errorf("failed to create edge db repository: %w", err)
	}

	return &EdgeDB{client: client}, nil
}

func (db *EdgeDB) Cleanup() error {
	return db.client.Close()
}
