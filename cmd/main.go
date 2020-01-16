package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anagram-service/config"
	"github.com/anagram-service/handler"
	log "github.com/sirupsen/logrus"
)

const appName = "anagram-service"

func main() {
	cfg, err := config.ParseConfig(appName)
	if err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	dict := make([]string, 0, 0)

	h, err := handler.CreateHTTPHandler(dict)
	if err != nil {
		log.Fatalf("creating http handler: %v", err)
	}

	listenErr := make(chan error, 1)
	server := &http.Server{
		Handler:      h,
		ReadTimeout:  cfg.API.ReadTimeout,
		WriteTimeout: cfg.API.WriteTimeout,
	}

	go func() {
		log.Println("ANAGRAM-SERVICE STARTED")
		listenErr <- http.ListenAndServe(cfg.API.Port, h)
	}()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-listenErr:
		log.Fatal(err)
	case <-osSignals:
		server.SetKeepAlivesEnabled(false)
		timeout := time.Second * 5

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
		log.Exit(0)
	}
}
