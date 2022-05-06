package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/taniko/rin/internal/app"
	"github.com/taniko/rin/internal/infrastructure/database"
	account "github.com/taniko/rin/internal/pb/taniko/rin/account/v1"
	channel "github.com/taniko/rin/internal/pb/taniko/rin/channel/v1"
	community "github.com/taniko/rin/internal/pb/taniko/rin/community/v1"
	message "github.com/taniko/rin/internal/pb/taniko/rin/message/v1"
	"github.com/taniko/rin/internal/server"
	"github.com/taniko/rin/internal/usecase"
	"github.com/taniko/sumire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	undefinedKeys := undefinedEnvs([]string{"APP_ENV", "PORT", "DB_USER", "DB_PASSWORD", "DB_HOST"})
	if len(undefinedKeys) > 0 {
		panic(fmt.Sprintf("undefined env: %s", strings.Join(undefinedKeys, ",")))
	}
	logger := sumire.NewLogger("sumire",
		sumire.WithStandardHandler(sumire.INFO, os.Stdout),
		sumire.WithRuntimeExtra(sumire.WARNING, sumire.DefaultRuntimeSkip),
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		logger.Critical(fmt.Sprintf("failed to listen: %v", err), nil)
		os.Exit(1)
	}
	s := grpc.NewServer()
	db, err := database.NewDatabase("account")
	if err != nil {
		logger.Critical(fmt.Sprintf("failed to connect database: %v", err), nil)
		os.Exit(1)
	}
	account.RegisterAccountServiceServer(s, server.NewAccountServer(
		logger,
		usecase.NewAccountUsecase(logger, database.NewAccountDatabase(db, logger)),
	))
	community.RegisterCommunityServiceServer(s, server.NewCommunityServer(logger))
	channel.RegisterChannelServiceServer(s, server.NewChannelServer(logger))
	message.RegisterMessageServiceServer(s, server.NewMessageServer(logger))

	if app.IsDevelopment() {
		reflection.Register(s)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		logger.Info("start", nil)
		if err := s.Serve(lis); err != nil && err != http.ErrServerClosed {
			logger.Critical(fmt.Sprintf("server error: %v", err), nil)
		}
	}()
	<-ctx.Done()
	s.GracefulStop()
	logger.Info("stop server", nil)
}

func undefinedEnvs(keys []string) []string {
	var result []string
	for _, key := range keys {
		if os.Getenv(key) == "" {
			result = append(result, key)
		}
	}
	return result
}
