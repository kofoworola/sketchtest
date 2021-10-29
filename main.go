package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port string `default:"8080"`
}

func main() {
	// configuration
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/draw", drawHandler)
	server := http.Server{
		Addr:        ":" + cfg.Port,
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
