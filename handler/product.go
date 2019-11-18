package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/sumitalp/productcatalog/models"
	"github.com/sumitalp/productcatalog/utils"
)

func (h *Handler) GetProduct(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.productStore.GetBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newProductResponse(c, a))
}

func (h *Handler) Products(c echo.Context) error {
	category := c.QueryParam("category")
	owner := c.QueryParam("owner")

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	var products []models.Product
	var count int
	if category != "" {
		products, count, err = h.productStore.ListByCategory(category, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else if owner != "" {
		products, count, err = h.productStore.ListByOwner(owner, offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	} else {
		products, count, err = h.productStore.List(offset, limit)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}
	}
	return c.JSON(http.StatusOK, newProductListResponse(h.userStore, userIDFromToken(c), products, count))
}

func (h *Handler) CreateProduct(c echo.Context) error {
	var a models.Product
	req := &productCreateRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	a.OwnerID = userIDFromToken(c)
	err := h.productStore.CreateProduct(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newProductResponse(c, &a))
}

func (h *Handler) UpdateProduct(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.productStore.GetUserProductBySlug(userIDFromToken(c), slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := &productUpdateRequest{}
	req.populate(a)
	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = h.productStore.UpdateProduct(a, req.Product.Categories); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newProductResponse(c, a))
}

func (h *Handler) DeleteProduct(c echo.Context) error {
	slug := c.Param("slug")
	a, err := h.productStore.GetUserProductBySlug(userIDFromToken(c), slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	err = h.productStore.DeleteProduct(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}

func (h *Handler) GetCategory(c echo.Context) error {
	categoryID64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(errors.New("Invalid ID.")))
	}
	categoryID := uint(categoryID64)
	a, err := h.productStore.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newCategoryResponse(c, a))
}

func (h *Handler) Categories(c echo.Context) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	var categories []models.Category
	var count int

	categories, count, err = h.productStore.ListCategories(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, newCategoryListResponse(h.userStore, userIDFromToken(c), categories, count))
}

func (h *Handler) CreateCategory(c echo.Context) error {
	var a models.Category
	req := &categoryCreateRequest{}
	if err := req.bind(c, &a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	err := h.productStore.CreateCategory(&a)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusCreated, newCategoryResponse(c, &a))
}

func (h *Handler) UpdateCategory(c echo.Context) error {
	categoryID64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(errors.New("Invalid ID.")))
	}
	categoryID := uint(categoryID64)
	a, err := h.productStore.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := &categoryUpdateRequest{}
	req.populate(a)
	if err := req.bind(c, a); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = h.productStore.UpdateCategory(a); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newCategoryResponse(c, a))
}

func (h *Handler) DeleteCategory(c echo.Context) error {
	categoryID64, err := strconv.ParseUint(c.Param("id"), 10, 64)
	
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewError(errors.New("Invalid ID.")))
	}
	categoryID := uint(categoryID64)
	a, err := h.productStore.GetCategoryByID(categoryID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if a == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	err = h.productStore.DeleteCategory(a)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"result": "ok"})
}