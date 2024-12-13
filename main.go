package main

import (
	"context"
	"embed"
	"log"

	"github.com/mgjules/mgjules-go/auth"
	"github.com/mgjules/mgjules-go/config"
	"github.com/mgjules/mgjules-go/fetcher"
	"github.com/mgjules/mgjules-go/http"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/projecter"
	"github.com/mgjules/mgjules-go/repository"
	"github.com/panjf2000/ants/v2"
)

//go:generate npm run build

//go:embed static/*
var static embed.FS

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	logger.Init(cfg.Prod)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := repository.New(ctx, cfg)

	pool, err := ants.NewPool(100)
	if err != nil {
		logger.L.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Release()

	fetcher := fetcher.New(repo, pool)
	fetcher.Start()
	defer fetcher.Stop()

	projecter := projecter.New(cfg.Prod, pool, fetcher)
	fetcher.AddSubscriber(projecter.Build)

	pool.Submit(fetcher.Fetch)

	auth := auth.New(cfg.AuthToken)

	server := http.NewServer(cfg.Prod,
		cfg.ServerHost,
		cfg.ServerPort,
		cfg.ServerTLSDomain,
		auth,
		fetcher,
		projecter,
		static,
	)
	if err = server.Start(); err != nil {
		logger.L.Fatalf("failed to start server: %v", err)
	}
}
