package service

import (
	"github.com/devfullcycle/imersao17/goapi/internal/database"
	"github.com/devfullcycle/imersao17/goapi/internal/entity"
)

type CategoryService struct {
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (categoryService *CategoryService) GetCategories() ([]*entity.Category, error) {
	categories, err := categoryService.CategoryDB.GetCategories()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (categoryService *CategoryService) CreateCategory(name string) (*entity.Category, error) {
	category := entity.NewCategory(name)

	_, err := categoryService.CategoryDB.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (categoryService *CategoryService) GetCategory(id string) (*entity.Category, error) {
	category, err := categoryService.CategoryDB.GetCategoryByID(id)

	if err != nil {
		return nil, err
	}

	return category, nil
}
