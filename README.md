# mgjules-go

A Go alternative of my [personal website](https://github.com/mgjules/mgjules). Gotta go fast!

## Development

Create a `.env` file with a valid `EDGEDB_DSN` and an `AUTH_TOKEN` of your choice:
```shell
EDGEDB_DSN="edgedb://edgedb:password@host:port/db"
AUTH_TOKEN="C7jI8iyCsaS9kiFRerVFpqvVzBedEEGVsFw1WSDN8mBQNxJ1dyi1qrWzKo8gOTbb0hmiK"
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

Create a `.env` file with a valid `PROD`, `EDGEDB_DSN` and an `AUTH_TOKEN` of your choice:
```shell
PROD=true
EDGEDB_DSN="edgedb://edgedb:password@host:port/db"
AUTH_TOKEN="C7jI8iyCsaS9kiFRerVFpqvVzBedEEGVsFw1WSDN8mBQNxJ1dyi1qrWzKo8gOTbb0hmiK"
```

```shell
$ go generate
$ go build .
```
