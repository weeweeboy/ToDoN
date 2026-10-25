package core_http_server

import (
	core_logger "ToDoN/internal/core/logger"
	core_http_middleware "ToDoN/internal/core/transport/http/middleware"
	"context"
	"errors"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type HTTPServer struct {
	mux         *http.ServeMux
	config      Config
	log         *core_logger.Logger
	middlewares []core_http_middleware.Middleware
}

func NewHTTPServer(config Config, log *core_logger.Logger, middlewares ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux:         http.NewServeMux(),
		config:      config,
		log:         log,
		middlewares: middlewares,
	}
}

func (s *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		h := core_http_middleware.ChainMiddlewares(router, s.middlewares...)
		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, h))
	}
}

func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddlewares(s.mux, s.middlewares...)

	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)

	go func() {
		defer close(ch)
		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	s.log.Warn("start http server", zap.String("addres", s.config.Addr))

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and server http: %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("shutdown http server...")
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer shutdownCancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown http server: %w", err)
		}
		s.log.Warn("shutdown http server successfully")
	}

	return nil
}
