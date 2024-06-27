package httpserver

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vaberof/smartway-task/pkg/http/httpserver/middleware/logging"
	"github.com/vaberof/smartway-task/pkg/logging/logs"
	"log/slog"
	"net/http"
)

type AppServer struct {
	Server    *http.Server
	ChiRouter *chi.Mux
	config    *ServerConfig
	logger    *slog.Logger
}

func New(config *ServerConfig, logsBuilder *logs.Logs) *AppServer {
	loggingMw := logging.New(logsBuilder)

	chiRouter := chi.NewRouter()

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
		Handler: loggingMw.Handler(chiRouter)}

	return &AppServer{
		Server:    httpServer,
		ChiRouter: chiRouter,
		config:    config,
		logger:    loggingMw.Logger,
	}
}

func (server *AppServer) StartAsync() <-chan error {
	exitChannel := make(chan error)

	go func() {
		err := server.Server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			exitChannel <- err
			return
		} else {
			exitChannel <- nil
		}
	}()

	return exitChannel
}
