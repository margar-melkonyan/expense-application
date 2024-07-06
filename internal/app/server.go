package app

import (
	"context"
	handler "expense-application/internal/handler"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) run(port string, handler *handler.Handler) error {
	slog.Info(fmt.Sprintf("Server started and listening on port %v", port))

	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf(":%s", port),
		Handler:        handler.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
