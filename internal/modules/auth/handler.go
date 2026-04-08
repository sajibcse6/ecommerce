package auth

import (
	"encoding/json"
	"net/http"

	"ecommerce/internal/pkg/response"
	"ecommerce/internal/pkg/validator"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	errs := validator.ValidationErrors{}
	validator.Required(input.Name, "name", errs)
	validator.Required(input.Email, "email", errs)
	validator.Email(input.Email, "email", errs)

	if len(errs) > 0 {
		response.Error(w, http.StatusBadRequest, errs)
		return
	}

	user, err := h.service.Register(r.Context(), input.Name, input.Email)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email string `json:"email"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	token,err := h.service.Login(r.Context(), input.Email)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{
		"token": token,
	})
}