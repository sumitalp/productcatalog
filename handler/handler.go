package handler

import (
	"github.com/sumitalp/productcatalog/product"
	"github.com/sumitalp/productcatalog/user"
)

type Handler struct {
	userStore    user.RepositoryInterface
	productStore product.RepositoryInterface
}

func NewHandler(ur user.RepositoryInterface, pr product.RepositoryInterface) *Handler {
	return &Handler{
		userStore:    ur,
		productStore: pr,
	}
}
