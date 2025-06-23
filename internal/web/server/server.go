package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pablisson/go-gateway/internal/service"
	"github.com/pablisson/go-gateway/internal/web/handlers"
)

type Server struct {
	router *chi.Mux
	server *http.Server
	accountService *service.AccountService
	port string
}

func NewServer(accountService *service.AccountService, port string) *Server {
	router := chi.NewRouter()

	return &Server{
		router: router,
		accountService: accountService,
		port: port,
	}
}

func (s *Server) ConfigureRoutes() {
	accountHandler := handlers.NewAccountHandler(s.accountService)
	//s.router.Post("/accounts", accountHandler.Create)
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/accounts", accountHandler.Get)
		r.Post("/accounts", accountHandler.Create)
	})
}

func (s *Server) Start() error {
	//s.ConfigureRoutes()

	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: s.router,
	}

	return s.server.ListenAndServe()
}
