package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/sumitalp/productcatalog/router/middleware"
	"github.com/sumitalp/productcatalog/utils"
)

func (h *Handler) Register(v1 *echo.Group) {
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	guestUsers := v1.Group("/users")
	guestUsers.POST("", h.SignUp)
	guestUsers.POST("/login", h.Login)

	user := v1.Group("/user", jwtMiddleware)
	user.GET("", h.CurrentUser)
	user.PUT("", h.UpdateUser)

	categories := v1.Group("categories", jwtMiddleware)
	categories.POST("", h.CreateCategory)
	categories.GET("", h.Categories)
	categories.GET("/:slug", h.GetCategory)
	categories.PUT("/:slug", h.UpdateCategory)
	categories.DELETE("/:slug", h.DeleteCategory)

	products := v1.Group("products", jwtMiddleware)
	products.POST("", h.CreateProduct)
	products.GET("", h.Products)
	products.GET("/:slug", h.GetProduct)
	products.PUT("/:slug", h.UpdateProduct)
	products.DELETE("/:slug", h.DeleteProduct)
}
