package handler

import (
	"time"
	"github.com/labstack/echo/v4"
	"github.com/sumitalp/productcatalog/models"
	"github.com/sumitalp/productcatalog/utils"
	"github.com/sumitalp/productcatalog/user"
)

type userResponse struct {
	User struct {
		Username string  `json:"username" xml:"username"`
		Email    string  `json:"email" xml:"email"`
		Bio      *string `json:"bio" xml:"bio"`
		Image    *string `json:"image" xml:"image"`
		Token    string  `json:"token" xml:"token"`
	} `json:"user" xml:"user"`
}

func newUserResponse(u *models.User) *userResponse {
	r := new(userResponse)
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Bio = u.Bio
	r.User.Image = u.Image
	r.User.Token = utils.GenerateJWT(u.ID)
	return r
}

type productResponse struct {
	Slug           string    `json:"slug" xml:"slug"`
	Title          string    `json:"title" xml:"title"`
	Description    string    `json:"description" xml:"description"`
	Image           string    `json:"image" xml:"image"`
	CategoryList        []string  `json:"categoryList" xml:"categories>category"`
	CreatedAt      time.Time `json:"createdAt" xml:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" xml:"updatedAt"`
	Owner         struct {
		Username  string  `json:"username" xml:"username"`
		Bio       *string `json:"bio" xml:"bio"`
		Image     *string `json:"image" xml:"image"`
	} `json:"owner" xml:"owner"`
}

type singleProductResponse struct {
	Product *productResponse `json:"product" xml:"product"`
}

type productListResponse struct {
	Products      []*productResponse `json:"products" xml:"products>product"`
	ProductsCount int                `json:"productsCount" xml:"productsCount"`
}

func newProductResponse(c echo.Context, a *models.Product) *singleProductResponse {
	ar := new(productResponse)
	ar.CategoryList = make([]string, 0)
	ar.Slug = a.Slug
	ar.Title = a.Title
	ar.Description = a.Description
	ar.Image = a.Image
	ar.CreatedAt = a.CreatedAt
	ar.UpdatedAt = a.UpdatedAt
	for _, t := range a.Categories {
		ar.CategoryList = append(ar.CategoryList, t.Category)
	}
	ar.Owner.Username = a.Owner.Username
	ar.Owner.Image = a.Owner.Image
	ar.Owner.Bio = a.Owner.Bio
	return &singleProductResponse{ar}
}

func newProductListResponse(us user.RepositoryInterface, userID uint, products []models.Product, count int) *productListResponse {
	r := new(productListResponse)
	r.Products = make([]*productResponse, 0)
	for _, a := range products {
		ar := new(productResponse)
		ar.CategoryList = make([]string, 0)
		ar.Slug = a.Slug
		ar.Title = a.Title
		ar.Description = a.Description
		ar.Image = a.Image
		ar.CreatedAt = a.CreatedAt
		ar.UpdatedAt = a.UpdatedAt
		for _, t := range a.Categories {
			ar.CategoryList = append(ar.CategoryList, t.Category)
		}

		ar.Owner.Username = a.Owner.Username
		ar.Owner.Image = a.Owner.Image
		ar.Owner.Bio = a.Owner.Bio

		r.Products = append(r.Products, ar)
	}
	r.ProductsCount = count
	return r
}

// Category
type categoryResponse struct {
	ID			   uint      `json:"id" xml:"id"`
	Title          string    `json:"title" xml:"title"`
	Description    string    `json:"description" xml:"description"`
	CreatedAt      time.Time `json:"createdAt" xml:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" xml:"updatedAt"`
}

type singleCategoryResponse struct {
	Category *categoryResponse `json:"category" xml:"category"`
}

type categoryListResponse struct {
	Categories      []*categoryResponse `json:"categories" xml:"categories>category"`
	CategoriesCount int                `json:"categoriesCount" xml:"categoriesCount"`
}

func newCategoryResponse(c echo.Context, a *models.Category) *singleCategoryResponse {
	ar := new(categoryResponse)
	ar.ID = a.ID
	ar.Title = a.Category
	ar.Description = a.Description
	ar.CreatedAt = a.CreatedAt
	ar.UpdatedAt = a.UpdatedAt

	return &singleCategoryResponse{ar}
}

func newCategoryListResponse(us user.RepositoryInterface, userID uint, categories []models.Category, count int) *categoryListResponse {
	r := new(categoryListResponse)
	r.Categories = make([]*categoryResponse, 0)
	for _, a := range categories {
		ar := new(categoryResponse)
		ar.ID = a.ID
		ar.Title = a.Category
		ar.Description = a.Description
		ar.CreatedAt = a.CreatedAt
		ar.UpdatedAt = a.UpdatedAt

		r.Categories = append(r.Categories, ar)
	}
	r.CategoriesCount = count
	return r
}

