package api

import (
	"sync"

	"github.com/ucok-man/P3-W1-PairProject/internal/config"
	"github.com/ucok-man/P3-W1-PairProject/internal/logging"
	"github.com/ucok-man/P3-W1-PairProject/internal/repo"
)

type Application struct {
	config *config.Config
	logger *logging.Logger
	repo   *repo.Services
	wg     sync.WaitGroup
	ctxkey struct {
		user string
	}
}

func New() *Application {
	logger := logging.New()
	cfg, err := config.New()
	if err != nil {
		logger.Fatal(err, "failed config initialization", nil)
	}

	dbconn, err := config.OpenDB(cfg)
	if err != nil {
		logger.Fatal(err, "failed open db connection", nil)
	}
	logger.Info("database connection pool established", nil)

	app := &Application{
		logger: logger,
		config: cfg,
		repo:   repo.New(dbconn),
		ctxkey: struct {
			user string
		}{
			user: "user",
		},
	}

	return app
}
