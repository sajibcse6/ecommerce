package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"ecommerce/internal/pkg/response"
	"ecommerce/internal/pkg/validator"
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
		response.Error(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// Validation
	errs := validator.ValidationErrors{}

	validator.Required(input.Name, "name", errs)
	validator.Required(input.Email, "email", errs)
	validator.Email(input.Email, "email", errs)

	if len(errs) > 0 {
		response.Error(w, http.StatusBadRequest, errs)
		return
	}

	user, err := h.service.CreateUser(r.Context(), input.Name, input.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, users)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)

	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetUserByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "User not found")
		return
	}

	response.JSON(w, http.StatusOK, user)
}