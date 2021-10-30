package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/kelseyhightower/envconfig"
	"github.com/kofoworola/sketchtest/storage/postgres"
)

type Config struct {
	Port string `default:"8080"`

	Postgres postgres.Config
}

func main() {
	// configuration
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	storage, err := postgres.New(&cfg.Postgres)
	if err != nil {
		log.Fatalf("error starting storage: %v", err)
	}
	if err := storage.Migrate(); err != nil && errors.Is(migrate.ErrNoChange, err) {
		log.Fatalf("error migrating: %v", err)
	}

	server := http.Server{
		Addr: ":" + cfg.Port,
		Handler: &Handler{
			store: storage,
		},
		ReadTimeout: time.Second * 10,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-done

		log.Println("server shutting down...")
		server.Shutdown(context.Background())
	}()

	log.Printf("starting server on port %s", cfg.Port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
