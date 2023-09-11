package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/joefazee/autodeploy/pkg/config"
	"github.com/joefazee/autodeploy/pkg/domain"
)

type server struct {
	store       domain.Store
	config      config.Config
	router      *gin.Engine
	asyncMux    *asynq.ServeMux
	asyncServer *asynq.Server
	asyncClient *asynq.Client
}

func newServer(store domain.Store, cfg config.Config) *server {
	return &server{
		store:  store,
		config: cfg,
	}
}

func (srv *server) run(address string) error {
	srv.setupRouter()

	if srv.router == nil {
		return domain.NewAppError("router instance cannot be nil", errors.New("in error"))
	}

	httpServer := &http.Server{
		Addr:         srv.config.HttServerAddress,
		Handler:      srv.router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)

	go func() {

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		log.Println("shutting down server ", s)

		// create a context with a 5-seconds timeout
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := httpServer.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		if srv.asyncServer != nil {
			srv.asyncServer.Shutdown()
		}

		shutdownError <- nil
	}()

	err := httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil

}
