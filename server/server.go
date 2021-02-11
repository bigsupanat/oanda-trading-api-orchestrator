package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/bigsupanat/oanda-trading-api-orchestrator/conf"
	"github.com/bigsupanat/oanda-trading-api-orchestrator/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server interface {
	Start()
	Stop() error
}

type server struct {
	router *chi.Mux
	srv    *http.Server
	svc    service.Service
}

func NewServer(svc service.Service) Server {
	router := chi.NewRouter()
	return &server{
		router: router,
		svc:    svc,
	}
}

func (s *server) Start() {
	go s.startServer()
}

func (s *server) startServer() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Get("/get_account_summary", getAccountSummaryFn(s.svc))

	s.router.Post("/createOrder", postOrderFn(s.svc))

	s.router.Put("/cancelOrder", cancelOrderFn(s.svc))

	s.srv = &http.Server{
		Addr:    conf.ServicePort,
		Handler: s.router,
	}
	if err := s.srv.ListenAndServe(); err != nil {
		log.Printf("listen: %s\n", err)
	}
}

func (s *server) Stop() error {
	if s.srv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.srv.Shutdown(ctx); err != nil {
			log.Println("Server Shutdown:", err)
			return err
		}
		return nil
	}
	return errors.New("No pointer to server")
}
