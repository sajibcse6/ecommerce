package server

import (
	"fmt"
	"net/http"

	"ecommerce/internal/config"
	"ecommerce/internal/modules/user"
	"ecommerce/internal/middleware"
	"ecommerce/internal/modules/auth"

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

	// Register middleware (Order Matters)
	r.Use(middleware.Recovery)
	r.Use(middleware.Logger)

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

	userRepo := user.NewRepository(s.db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	//auth
	authService := auth.NewService(userRepo, s.config.JWTSecret)
	authHandler := auth.NewHandler(authService)
	auth.RegisterRoutes(s.router, authHandler)

	// protected routes
	s.router.Group(func (r chi.Router) {
		r.Use(middleware.Auth(s.config.JWTSecret))

		user.RegisterRoutes(r, userHandler)
	})
}

func (s *Server) Start() error {
	fmt.Println("Server running on port", s.config.Port)
	return http.ListenAndServe("127.0.0.1:"+s.config.Port, s.router)
}
