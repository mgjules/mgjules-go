package main

import (
	"context"
	"log"

	"github.com/mgjules/mgjules-go/internal/auth"
	"github.com/mgjules/mgjules-go/internal/config"
	"github.com/mgjules/mgjules-go/internal/fetcher"
	"github.com/mgjules/mgjules-go/internal/http"
	"github.com/mgjules/mgjules-go/internal/logger"
	"github.com/mgjules/mgjules-go/internal/projecter"
	"github.com/mgjules/mgjules-go/internal/repository"
	"github.com/panjf2000/ants/v2"
)

//go:generate npm run build

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}

	logger.Init(cfg.Prod)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := repository.New(ctx, cfg)
	defer repo.Cleanup()

	pool, err := ants.NewPool(100)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Release()

	fetcher := fetcher.New(repo, pool)
	fetcher.Start()
	defer fetcher.Stop()

	projecter := projecter.New(cfg.Prod, pool, fetcher)
	fetcher.AddSubscriber(projecter.Build)

	pool.Submit(fetcher.Fetch)

	auth := auth.New(cfg.AuthToken)

	server := http.NewServer(
		cfg.Prod,
		cfg.ServerHost,
		cfg.ServerPort,
		cfg.ServerTLSDomain,
		auth,
		fetcher,
		projecter,
	)
	if err = server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
