package main

import (
	core_logger "ToDoN/internal/core/logger"
	core_postgres_pool "ToDoN/internal/core/repository/postgres/pool"
	core_http_middleware "ToDoN/internal/core/transport/http/middleware"
	core_http_server "ToDoN/internal/core/transport/http/server"
	users_postgres_repository "ToDoN/internal/features/users/repository/postgres"
	users_service "ToDoN/internal/features/users/service"
	users_trnsport_http "ToDoN/internal/features/users/transport/http"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to initialize logger")
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Start ToDo Application!")

	logger.Debug("initialising connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(core_postgres_pool.NewConfigMust(), ctx)
	if err != nil {
		logger.Fatal("Failed to initialize postgres pool", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("initialising feature", zap.String("feature_name", "users"))

	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUserService(usersRepository)
	usersTransportHTTP := users_trnsport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initialising HTTP server")

	httpServer := core_http_server.NewHTTPServer(core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace())

	apiVersionRouter := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("Failed to start HTTP server", zap.Error(err))
	}

}
