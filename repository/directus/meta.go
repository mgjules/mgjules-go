package directus

import (
	"context"
	"errors"

	"github.com/mgjules/mgjules-go/entity"
)

func (db *Directus) GetMeta(ctx context.Context, id string) (*entity.Meta, error) {
	return nil, errors.New("not implemented")
}
