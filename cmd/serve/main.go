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

	repo, err := repository.New(ctx, cfg)
	if err != nil {
		log.Fatalf("failed to create repository: %v", err)
	}
	defer repo.Cleanup()

	pool, err := ants.NewPool(100)
	if err != nil {
		log.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Release()

	f, err := fetcher.New(repo, pool)
	if err != nil {
		log.Fatalf("failed to create fetcher: %v", err)
	}
	f.Start()
	defer f.Stop()

	p, err := projecter.New(cfg.Prod, pool, f)
	if err != nil {
		log.Fatalf("failed to create projecter: %v", err)
	}
	f.AddSubscriber(p.Build)

	pool.Submit(f.Fetch)

	a, err := auth.New(cfg.AuthToken)
	if err != nil {
		log.Fatalf("failed to create auth: %v", err)
	}

	server, err := http.NewServer(
		cfg.Prod,
		cfg.ServerHost,
		cfg.ServerPort,
		cfg.ServerTLSDomain,
		a,
		f,
		p,
	)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}
	if err = server.Start(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
