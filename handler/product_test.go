package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/sumitalp/productcatalog/router"
	"github.com/sumitalp/productcatalog/router/middleware"
	"github.com/sumitalp/productcatalog/utils"
)


// Product Test cases
func TestListProductsCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	e := router.New()
	req := httptest.NewRequest(echo.GET, "/api/products", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.Products(c)
	})(c)
	
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var aa productListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &aa)
		assert.NoError(t, err)
		assert.NotNil(t, &aa)
	}
}

func TestGetProductCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.GET, "/api/products/:slug", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/products/:slug")
	c.SetParamNames("slug")
	c.SetParamValues("product1-slug")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.GetProduct(c)
	})(c)
	assert.NoError(t, err)
	
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var a singleProductResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "product1-slug", a.Product.Slug)
	}
}

func TestCreateProductCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"product":{"title":"product2", "description":"product2",  "categoryList":["category1","category2"]}}`
	)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.POST, "/api/products", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.CreateProduct(c)
	})(c)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var a singleProductResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "product2", a.Product.Slug)
		assert.Equal(t, "product2", a.Product.Description)
		assert.Equal(t, "product2", a.Product.Title)
		assert.Equal(t, "user1", a.Product.Owner.Username)
		assert.Equal(t, 2, len(a.Product.CategoryList))
	}
}

func TestUpdateProductCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"product":{"title":"product1 part 2", "categoryList":["category3"]}}`
	)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.PUT, "/api/products/:slug", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/products/:slug")
	c.SetParamNames("slug")
	c.SetParamValues("product1-slug")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.UpdateProduct(c)
	})(c)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, rec.Code) {
		var a singleProductResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "product1 part 2", a.Product.Title)
		assert.Equal(t, "product1-part-2", a.Product.Slug)
	}
}

func TestDeleteProductCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.DELETE, "/api/products/:slug", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/products/:slug")
	c.SetParamNames("slug")
	c.SetParamValues("product1-slug")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.DeleteProduct(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}


// Category Test cases
func TestListCategoriesCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	e := router.New()
	req := httptest.NewRequest(echo.GET, "/api/categories", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.Categories(c)
	})(c)
	
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var aa categoryListResponse
		err := json.Unmarshal(rec.Body.Bytes(), &aa)
		assert.NoError(t, err)
		assert.NotNil(t, &aa)
	}
}

func TestGetCategoryCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.GET, "/api/categories/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/products/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.GetCategory(c)
	})(c)
	assert.NoError(t, err)
	
	if assert.Equal(t, http.StatusOK, rec.Code) {
		var a singleCategoryResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "category1", a.Category.Title)
	}
}

func TestCreateCategoryCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"category":{"title":"category3", "description":"category3"}}`
	)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.POST, "/api/categories", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return h.CreateCategory(c)
	})(c)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusCreated, rec.Code) {
		var a singleCategoryResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "category3", a.Category.Description)
		assert.Equal(t, "category3", a.Category.Title)
	}
}

func TestUpdateCategoryCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"category":{"title":"category1 part 2"}}`
	)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.PUT, "/api/categories/:id", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/categories/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.UpdateCategory(c)
	})(c)
	assert.NoError(t, err)

	if assert.Equal(t, http.StatusOK, rec.Code) {
		var a singleCategoryResponse
		err := json.Unmarshal(rec.Body.Bytes(), &a)
		assert.NoError(t, err)
		assert.Equal(t, "category1 part 2", a.Category.Title)
	}
}

func TestDeleteCategoryCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.DELETE, "/api/categories/:id", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, authHeader(utils.GenerateJWT(1)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/categories/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	err := jwtMiddleware(func(context echo.Context) error {
		return h.DeleteCategory(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
