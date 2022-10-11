package main

import (
	"context"
	"embed"
	"log"

	"github.com/mgjules/mgjules-go/auth"
	"github.com/mgjules/mgjules-go/config"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/projection"
	"github.com/mgjules/mgjules-go/repository/edgedb"
	"github.com/mgjules/mgjules-go/server"
)

//go:generate npm run build

//go:embed templates
var templates embed.FS

//go:embed static
var static embed.FS

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	logger.Init(cfg.Prod)

	ctx := context.Background()
	repo, err := edgedb.New(ctx, cfg.EdgeDBDSN)
	if err != nil {
		logger.Logger.Fatalf("failed to create edgedb repository: %v", err)
	}

	auth := auth.New(cfg.AuthToken)

	projection, err := projection.New(cfg.Prod, repo, templates)
	if err != nil {
		logger.Logger.Fatalf("failed to create projection: %v", err)
	}

	projection.Start()
	projection.FetchData()
	projection.BuildProjections()

	server := server.New(cfg.Prod, cfg.ServerHost, cfg.ServerPort, auth, projection, static)
	if err = server.Start(); err != nil {
		logger.Logger.Fatalf("failed to start server: %v", err)
	}
}
