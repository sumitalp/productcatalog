package product

import (
	"github.com/sumitalp/productcatalog/models"
)

type RepositoryInterface interface {
	GetBySlug(string) (*models.Product, error)
	GetUserProductBySlug(userID uint, slug string) (*models.Product, error)
	CreateProduct(*models.Product) error
	UpdateProduct(*models.Product, []string) error
	DeleteProduct(*models.Product) error
	List(offset, limit int) ([]models.Product, int, error)
	ListByCategory(category string, offset, limit int) ([]models.Product, int, error)
	ListByOwner(username string, offset, limit int) ([]models.Product, int, error)

	ListCategories(offset, limit int) ([]models.Category, int, error)
	CreateCategory(*models.Category) error
	UpdateCategory(*models.Category) error
	DeleteCategory(*models.Category) error
	GetCategoryByID(uint) (*models.Category, error)
}
