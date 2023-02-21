package server

import (
	"context"
	"github.com/sirupsen/logrus"
	"help-service/api"
	"help-service/internals/app/controller"
	"help-service/internals/config"
	"net/http"
	"time"
)

type Server struct {
	cfg config.Config
	srv *http.Server
}

var log = logrus.New()

func New(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (server *Server) Start() {
	log.Infoln("Server starting")

	helpController := controller.New()

	route := api.CreateRoute(helpController)

	server.srv = &http.Server{
		Addr:    ":" + server.cfg.Port,
		Handler: route,
	}

	log.Infoln("Server started")

	err := server.srv.ListenAndServe()
	if err != nil {
		log.Fatalln("Error listen server: ", err)
	}
}

func (server *Server) Shutdown() {
	log.Infoln("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()
	}()

	var err error

	if err = server.srv.Shutdown(ctx); err != nil {
		log.Fatalln("Server shutdown failed: ", err)
	}

	if err == http.ErrServerClosed {
		err = nil
	}
}
