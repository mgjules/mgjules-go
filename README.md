# mgjules-go

[![Go version](https://img.shields.io/github/go-mod/go-version/mgjules/mgjules-go.svg)](https://pkg.go.dev/github.com/mgjules/mgjules-go)

Source code for my [personal website](https://michael.jules.mu).

## Architecture

> For the curious. Yes, it's over-engineered on purpose — this is a playground for experimentation.

All pages are pre-rendered into memory at startup and refreshed every hour. HTTP requests are plain memory lookups — no per-request templating.

```
Repository (Directus / Static)
    └── Fetcher  — concurrent fetch, hourly cron, in-memory cache
          └── Projecter  — builds pre-rendered HTML pages on each fetch
                └── HTTP Server  — serves pages directly from memory
```

In development mode, pages rebuild on every request so template changes are visible immediately without restarting.

## Prerequisites

- **Go 1.22+**
- **Node.js** (LTS recommended) — only needed to build CSS via [UnoCSS](https://unocss.dev)

## Configuration

All config is via environment variables. Copy `.env.example` to get started:

```shell
cp .env.example .env
```

| Variable | Required | Default | Description |
|---|---|---|---|
| `PROD` | | `false` | Enables JSON logging, INFO level, and TLS |
| `SERVER_HOST` | | `0.0.0.0` | Listen address |
| `SERVER_PORT` | | `13337` | Listen port |
| `SERVER_TLS_DOMAIN` | | | Domain for automatic TLS via Let's Encrypt |
| `AUTH_TOKEN` | yes | | Token for future protected endpoints |
| `STATIC` | | `false` | Use hardcoded static data (no CMS needed) |
| `DIRECTUS_URL` | | | Directus CMS base URL |
| `DIRECTUS_TOKEN` | | | Directus API token |
**Repository selection priority:** Directus → Static. At least one must be configured.

**Content:** lives in the configured CMS (Directus). For local development without CMS access, set `STATIC=true` to use the hardcoded data in `internal/repository/static/`.

## Development

```shell
# Install Node deps (first time only)
npm install

# In one terminal — watch and rebuild CSS
npm run dev

# In another terminal — run the server
cp .env.example .env  # then edit as needed
go run ./cmd/serve
```

## Production

`go generate` compiles the CSS (`npm run build`) before the binary is built — make sure Node deps are installed first.

```shell
npm install
go generate
go build -o mgjules-go ./cmd/serve
./mgjules-go
```

Sample production `.env`:

```shell
PROD=true
SERVER_HOST=0.0.0.0
SERVER_PORT=80
SERVER_TLS_DOMAIN=michael.jules.mu
AUTH_TOKEN=a_very_long_secret_token
DIRECTUS_URL=https://directus.example.com
DIRECTUS_TOKEN=a_very_long_secret_token
```
