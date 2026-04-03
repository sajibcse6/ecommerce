package server

import (
	"fmt"
	"net/http"

	"ecommerce/internal/config"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	router *chi.Mux
	config *config.Config
	db     *pgxpool.Pool
}

func New(cfg *config.Config, db *pgxpool.Pool) *Server {
	r := chi.NewRouter()

	s := &Server{
		router: r,
		config: cfg,
		db:     db,
	}

	s.registerRoutes()

	return s
}

func (s *Server) registerRoutes() {
	s.router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func (s *Server) Start() error {
	fmt.Println("Server running on port", s.config.Port)
	return http.ListenAndServe("127.0.0.1:"+s.config.Port, s.router)
}
