package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/taniko/rin/internal/app"
	"github.com/taniko/rin/internal/app/env"
	"github.com/taniko/rin/internal/gen/taniko/rin/account/v1/accountv1connect"
	"github.com/taniko/rin/internal/handler"
	"github.com/taniko/rin/internal/infrastructure/database"
	"github.com/taniko/rin/internal/infrastructure/local"
	"github.com/taniko/rin/internal/usecase"
)

func main() {
	ctx := context.Background()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	e, err := env.Load()
	if err != nil {
		panic(err)
	} else if e.IsNull() {
		panic("environment is null")
	}
	environment := e.Value()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if app.IsDevelopment() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	client, err := local.NewSecretClient(ctx)
	if err != nil {
		log.Panic().Err(err)
	}
	dbPassword, err := client.GetSecret(ctx, environment.DB.Key)
	if err != nil {
		log.Panic().Err(err)
	}
	db, err := database.New(environment, database.Password(dbPassword))
	if err != nil {
		panic(err)
	}

	accountRepo := database.NewAccount(db)

	signingKey, err := client.GetSecret(ctx, "signing-key")
	if err != nil {
		log.Panic().Err(err)
	}
	mux := http.NewServeMux()
	mux.Handle(accountv1connect.NewAccountServiceHandler(handler.NewAccount(usecase.NewAccountUsecase(
		accountRepo,
		[]byte(signingKey)),
	)))
	srv := &http.Server{
		Addr:              environment.Host,
		Handler:           mux,
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024,
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		log.Info().Msg("start server")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("server listen")
		}
	}()
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server shutdown")
	}
}
