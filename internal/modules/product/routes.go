package product

import (
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r chi.Router, h *Handler) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/", h.GetProducts)
		r.Get("/{id}", h.GetProductByID)
	})
}
