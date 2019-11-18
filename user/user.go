package user

import (
	"github.com/sumitalp/productcatalog/models"
)

type RepositoryInterface interface {
	GetByID(uint) (*models.User, error)
	GetByEmail(string) (*models.User, error)
	GetByUsername(string) (*models.User, error)
	Create(*models.User) error
	Update(*models.User) error
}
