package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

const appName = "anagram-service"

func main() {
	config, err := parseConfig(appName)
	if err != nil {
		log.Fatalf("parsing config: %v", err)
	}

	dict := make([]string, 0, 0)

	handler, err := createHTTPHandler(dict)
	if err != nil {
		log.Fatalf("creating http handler: %v", err)
	}

	listenErr := make(chan error, 1)
	server := &http.Server{
		ReadTimeout:  config.API.ReadTimeout,
		WriteTimeout: config.API.WriteTimeout,
	}

	go func() {
		log.Println("ANAGRAM-SERVICE STARTED")
		listenErr <- http.ListenAndServe(config.API.Port, handler)
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
