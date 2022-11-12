package main

import (
	"context"
	"embed"
	"log"

	"github.com/bep/godartsass"
	"github.com/mgjules/mgjules-go/auth"
	"github.com/mgjules/mgjules-go/config"
	"github.com/mgjules/mgjules-go/fetcher"
	"github.com/mgjules/mgjules-go/http"
	"github.com/mgjules/mgjules-go/logger"
	"github.com/mgjules/mgjules-go/projecter"
	"github.com/mgjules/mgjules-go/repository"
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"
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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	pool, err := ants.NewPool(1000)
	if err != nil {
		logger.L.Fatalf("failed to create pool: %v", err)
	}
	defer pool.Release()

	fetcher := fetcher.New(repo, pool, cron.New())
	fetcher.Start()
	go fetcher.Fetch(ctx)
	defer fetcher.Stop()

	projecter, err := projecter.New(cfg.Prod, pool, fetcher, templates, transpiler)
	if err != nil {
		logger.L.Fatalf("failed to create projecter: %v", err)
	}
	fetcher.AddSubscriber(projecter.Build)

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
