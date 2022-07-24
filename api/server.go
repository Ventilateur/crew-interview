package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ventilateur/crew-interview/config"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg config.ServerConfig, handler HTTPHandler) *Server {
	// Default engine Logger and Recovery middleware already attached
	router := gin.Default()

	router.GET("/health", handler.Health)

	v1Talents := router.Group("/v1/talents")
	v1Talents.GET("", handler.ListTalents)
	v1Talents.POST("", handler.AddTalent)

	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: router,
		},
	}
}

func (s *Server) Serve() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Initializing the server in a goroutine so that it won't block the graceful shutdown handling below
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.WithError(err).Fatal("error starting http server")
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// The context is used to inform the server it has x seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(60)*time.Second,
	)
	defer cancel()

	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.WithError(err).Fatal("error shutting http server down")
	}
}
