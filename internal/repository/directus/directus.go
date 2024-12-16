package directus

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

type directus struct {
	directusURL string
	client      *resty.Client
}

func New(ctx context.Context, isProd bool, url, token string) *directus {
	client := resty.New()

	if !isProd {
		client.SetDebug(true)
	}

	client.SetAuthToken(token)
	client.SetBaseURL(url)
	client.SetTimeout(30 * time.Second)

	return &directus{
		directusURL: url,
		client:      client,
	}
}

func (*directus) Cleanup() error {
	return nil
}
