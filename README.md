# mgjules-go

[![Go version](https://img.shields.io/github/go-mod/go-version/mgjules/mgjules-go.svg)](https://pkg.go.dev/github.com/mgjules/mgjules-go)
[![Uptime](https://status.mgjules.dev/api/v1/endpoints/core_myspace/uptimes/7d/badge.svg)](https://status.mgjules.dev/endpoints/core_myspace)
[![Uptime](https://status.mgjules.dev/api/v1/endpoints/core_myspace/response-times/7d/badge.svg)](https://status.mgjules.dev/endpoints/core_myspace)

Source code for my [personal website](https://mgjules.dev) hosted on a raspberry pi 3b+.

## Background

This website has seen several iterations through the years and moved across many repositories.

This iteration has for primary goal to experiment with the idea of pre-rendering the various views and serving them from memory. Obviously it is over-engineered but that's the whole point; to experiment without limits :wink:

## Development

Sample `.env`:
```shell
PROD=false
SERVER_PORT=13337
AUTH_TOKEN=a_very_long_but_obviously_fake_token
EDGEDB_DSN=edgedb://edgedb:fakepassword@edgy.mgjules.dev/myspace
DIRECTUS_URL=https://directus.mgjules.dev
DIRECTUS_TOKEN=a_very_long_but_obviously_fake_token
```

Run `go` app:
```shell
$ go run .
```

Run `unocss-cli` in `watch` mode:
```shell
$ npm run dev
```

## Production

Sample `.env`:
```shell
PROD=true
SERVER_PORT=80
AUTH_TOKEN=a_very_long_but_obviously_fake_token
EDGEDB_DSN=edgedb://edgedb:fakepassword@edgy.mgjules.dev/myspace
DIRECTUS_URL=https://directus.mgjules.dev
DIRECTUS_TOKEN=a_very_long_but_obviously_fake_token
```

```shell
$ go generate
$ go build .
```
