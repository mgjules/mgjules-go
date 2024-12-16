package edgedb

import (
	"context"
	"fmt"

	"github.com/edgedb/edgedb-go"
)

type edgeDB struct {
	client *edgedb.Client
}

func New(ctx context.Context, dsn string) (*edgeDB, error) {
	client, err := edgedb.CreateClientDSN(ctx, dsn, edgedb.Options{})
	if err != nil {
		return nil, fmt.Errorf("failed to create edge db repository: %w", err)
	}

	return &edgeDB{client: client}, nil
}

func (db *edgeDB) Cleanup() error {
	return db.client.Close()
}
