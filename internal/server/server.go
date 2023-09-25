package server

import (
	"FreeMusic/internal/config"
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

type server struct {
	httpServer *http.Server
	config     config.Config
}

func NewServer(conf *config.Config) *server {
	return &server{
		httpServer: nil,
		config:     *conf,
	}
}

func (s *server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", s.config.AppHost, s.config.AppPort),
		Handler:        handler,
		MaxHeaderBytes: s.config.AppMaxHeaderBytes, // 1 MB
		ReadTimeout:    time.Duration(s.config.AppReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(s.config.AppWriteTimeout) * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
