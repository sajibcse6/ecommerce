package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}

	err := json.NewDecoder(r.Body).Decode(&input);
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(r.Context(), input.Name, input.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idParam, 10, 64)

	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}