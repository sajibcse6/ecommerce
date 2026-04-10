package product

import (
	"ecommerce/internal/pkg/response"
	"ecommerce/internal/pkg/validator"
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

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
	}

	json.NewDecoder(r.Body).Decode(&input)

	errs := validator.ValidationErrors{}
	validator.Required(input.Name, "name", errs)

	if input.Price <= 0 {
		errs["price"] = "price must be greater than 0"
	}

	if len(errs) > 0 {
		response.Error(w, http.StatusBadRequest, errs)
		return
	}

	product, err := h.service.CreateProduct(r.Context(), input.Name, input.Description, input.Price)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, product)
}

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetProducts(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, products)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	product, err := h.service.GetProductByID(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "Product not found")
		return
	}

	response.JSON(w, http.StatusOK, product)
}
