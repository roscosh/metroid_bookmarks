package misc

import (
	"context"
	"net/http"
	"time"
)

const ReadHeaderTimeout = 20 * time.Second

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:              ":3000",
		Handler:           handler,
		ReadHeaderTimeout: ReadHeaderTimeout,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
