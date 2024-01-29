package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/imersao17/goapi/internal/entity"
	"github.com/devfullcycle/imersao17/goapi/internal/service"
	"github.com/go-chi/chi"
)

type WebCategoryHandler struct {
	CategoryService *service.CategoryService
}

func NewWebCategoryHandler(categoryService *service.CategoryService) *WebCategoryHandler {
	return &WebCategoryHandler{CategoryService: categoryService}
}

func (webCategoryHandler *WebCategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := webCategoryHandler.CategoryService.GetCategories()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(categories)
}

func (webCategoryHandler *WebCategoryHandler) GetCategory(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	category, err := webCategoryHandler.CategoryService.GetCategory(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(category)
}

func (webCategoryHandler *WebCategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.Category

	err := json.NewDecoder(r.Body).Decode(&category)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := webCategoryHandler.CategoryService.CreateCategory(category.Name)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
