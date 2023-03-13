package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"

	"github.com/sirjager/go_emails/cfg"

	"github.com/sirjager/go_emails/cmd/server"

	emailsServer "github.com/sirjager/go_emails/pkg/server"
)

var logger zerolog.Logger
var startTime time.Time

func init() {
	startTime = time.Now()
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	logger = logger.With().Timestamp().Logger()
	logger = logger.With().Str("service", "emails").Logger()
}

func main() {
	config, err := cfg.LoadConfig()
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to load configs")
	}

	srvr, err := emailsServer.NewEmailsServer(startTime, logger, config)
	if err != nil {
		logger.Fatal().Err(err).Msg("unable to create emails service")
	}

	errs := make(chan error)

	go handleSignals(errs)
	go server.RunGRPCServer(srvr, logger, config, errs)

	if config.Http != "" {
		// Made this optional since this service will only be used by internal services
		go server.RunGatewayServer(srvr, logger, config, errs)
	}

	logger.Error().Err(<-errs).Msg("exit")
}

func handleSignals(errs chan error) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	errs <- fmt.Errorf("%s", <-c)
}
