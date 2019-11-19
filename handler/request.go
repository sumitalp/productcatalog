package handler

import (
	"github.com/gosimple/slug"
	"github.com/labstack/echo/v4"
	"github.com/sumitalp/productcatalog/models"
)

type userRegisterRequest struct {
	User struct {
		Username string `json:"username" validate:"required" xml:"username"`
		Email    string `json:"email" validate:"required,email" xml:"email"`
		Password string `json:"password" validate:"required" xml:"password"`
	} `json:"user" xml:"user"`
}

func (r *userRegisterRequest) bind(c echo.Context, u *models.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	h, err := u.HashPassword(r.User.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}

type userUpdateRequest struct {
	User struct {
		Username string `json:"username" xml:"username"`
		Email    string `json:"email" validate:"email" xml:"email"`
		Password string `json:"password" xml:"password"`
		Bio      string `json:"bio" xml:"bio"`
		Image    string `json:"image" xml:"image"`
	} `json:"user" xml:"user"`
}

func newUserUpdateRequest() *userUpdateRequest {
	return new(userUpdateRequest)
}

func (r *userUpdateRequest) populate(u *models.User) {
	r.User.Username = u.Username
	r.User.Email = u.Email
	r.User.Password = u.Password
	if u.Bio != nil {
		r.User.Bio = *u.Bio
	}
	if u.Image != nil {
		r.User.Image = *u.Image
	}
}

func (r *userUpdateRequest) bind(c echo.Context, u *models.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.User.Username
	u.Email = r.User.Email
	if r.User.Password != u.Password {
		h, err := u.HashPassword(r.User.Password)
		if err != nil {
			return err
		}
		u.Password = h
	}
	u.Bio = &r.User.Bio
	u.Image = &r.User.Image
	return nil
}

type userLoginRequest struct {
	User struct {
		Email    string `json:"email" validate:"required,email" xml:"email"`
		Password string `json:"password" validate:"required" xml:"password"`
	} `json:"user" xml:"user"`
}

func (r *userLoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

// Product
type productCreateRequest struct {
	Product struct {
		Title       string   `json:"title" validate:"required" xml:"title"`
		Description string   `json:"description" validate:"required" xml:"description"`
		Image        string   `json:"image" xml:"image"`
		Categories        []string `json:"categoryList, omitempty" xml:"categories>category"`
	} `json:"product" xml:"product"`
}

func (r *productCreateRequest) bind(c echo.Context, a *models.Product) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Product.Title
	a.Slug = slug.Make(r.Product.Title)
	a.Description = r.Product.Description
	a.Image = r.Product.Image
	if r.Product.Categories != nil {
		for _, t := range r.Product.Categories {
			a.Categories = append(a.Categories, models.Category{Category: t})
		}
	}
	return nil
}

type productUpdateRequest struct {
	Product struct {
		Title       string   `json:"title" xml:"title"`
		Description string   `json:"description" xml:"description"`
		Image        string   `json:"image" xml:"image"`
		Categories        []string `json:"categoriesList" xml:"categories>category"`
	} `json:"product" xml:"product"`
}

func (r *productUpdateRequest) populate(a *models.Product) {
	r.Product.Title = a.Title
	r.Product.Description = a.Description
}

func (r *productUpdateRequest) bind(c echo.Context, a *models.Product) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Title = r.Product.Title
	a.Slug = slug.Make(a.Title)
	a.Description = r.Product.Description
	a.Image = r.Product.Image
	return nil
}

// Category
type categoryCreateRequest struct {
	Category struct {
		Title       string   `json:"title" validate:"required" xml:"title"`
		Description string   `json:"description" xml:"description"`
	} `json:"category" xml:"category"`
}

func (r *categoryCreateRequest) bind(c echo.Context, a *models.Category) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Category = r.Category.Title
	a.Description = r.Category.Description

	return nil
}

type categoryUpdateRequest struct {
	Category struct {
		Title       string   `json:"title" xml:"title"`
		Description string   `json:"description" xml:"description"`
	} `json:"category" xml:"category"`
}

func (r *categoryUpdateRequest) populate(c *models.Category) {
	r.Category.Title = c.Category
	r.Category.Description = c.Description
}

func (r *categoryUpdateRequest) bind(c echo.Context, a *models.Category) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	a.Category = r.Category.Title
	a.Description = r.Category.Description
	return nil
}