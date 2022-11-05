package main

import (
	"context"
	"embed"
	"log"

	"github.com/bep/godartsass"
	"github.com/mgjules/mgjules-go/auth"
	"github.com/mgjules/mgjules-go/config"
	"github.com/mgjules/mgjules-go/http"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/projection"
	"github.com/mgjules/mgjules-go/repository"
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
	repo, err := repository.New(ctx, cfg)
	if err != nil {
		logger.L.Fatalf("failed to create repository: %v", err)
	}

	auth := auth.New(cfg.AuthToken)

	transpiler, err := godartsass.Start(godartsass.Options{
		DartSassEmbeddedFilename: cfg.DartSassEmbeddedBinary,
	})
	if err != nil {
		logger.L.Fatalf("failed to start transpiler: %v", err)
	}
	defer transpiler.Close()

	projection, err := projection.New(cfg.Prod, repo, templates, transpiler)
	if err != nil {
		logger.L.Fatalf("failed to create projection: %v", err)
	}

	projection.Start()
	projection.FetchData()
	projection.BuildProjections()

	server := http.NewServer(cfg.Prod, cfg.ServerHost, cfg.ServerPort, cfg.ServerTLSDomain, auth, projection, static)
	if err = server.Start(); err != nil {
		logger.L.Fatalf("failed to start server: %v", err)
	}
}
