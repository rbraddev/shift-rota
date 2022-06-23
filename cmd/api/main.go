package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/rbraddev/shift-rota/internal/database"
	"github.com/rbraddev/shift-rota/internal/log"
	"github.com/rbraddev/shift-rota/internal/version"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn         string
		automigrate bool
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
	version bool
}

type application struct {
	config config
	db     *database.DB
	logger *log.Logger
	wg     sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "postgres DSN")
	flag.BoolVar(&cfg.db.automigrate, "db-automigrate", true, "run migrations on startup")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.BoolVar(&cfg.version, "version", false, "display version and exit")

	flag.Parse()

	if cfg.version {
		fmt.Printf("version: %s\n", version.Get())
		return
	}

	logger := log.New(os.Stdout, log.LevelAll)

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	logger.Info("database connection pool established", nil)

	app := &application{
		config: cfg,
		db:     db,
		logger: logger,
	}

	err = app.Run(cfg.port, app.routes())
	if err != nil {
		logger.Fatal(err)
	}

}
