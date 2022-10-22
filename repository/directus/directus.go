package directus

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

type Directus struct {
	client *resty.Client
}

func New(ctx context.Context, isProd bool, url, token string) *Directus {
	client := resty.New()

	if !isProd {
		client.SetDebug(true)
	}

	client.SetAuthToken(token)
	client.SetBaseURL(url)
	client.SetTimeout(30 * time.Second)

	return &Directus{client: client}
}
