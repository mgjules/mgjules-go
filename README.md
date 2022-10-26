# mgjules-go

Source code for my [personal website](https://github.com/mgjules/mgjules).

## Background

This website has seen several iterations through the years and moved across many repositories.

This iteration has for primary goal to experiment with the idea of pre-rendering the various views and serving them from memory. This obviously is over-engineered but that's the whole point; to experiment without limits :wink:

## Development

Create a `.env` file with a valid `EDGEDB_DSN` and an `AUTH_TOKEN` of your choice:
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

Create a `.env` file with a valid `PROD`, `EDGEDB_DSN`, `DIRECTUS_URL`, `DIRECTUS_TOKEN`, and an `AUTH_TOKEN` of your choice:
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
