package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sumitalp/productcatalog/models"
	"github.com/sumitalp/productcatalog/utils"
)

func (h *Handler) SignUp(c echo.Context) error {
	var u models.User
	req := &userRegisterRequest{}
	if err := req.bind(c, &u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userStore.Create(&u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newUserResponse(&u))
}

func (h *Handler) Login(c echo.Context) error {
	req := &userLoginRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	u, err := h.userStore.GetByEmail(req.User.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	if !u.CheckPassword(req.User.Password) {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}

func (h *Handler) CurrentUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}

func (h *Handler) UpdateUser(c echo.Context) error {
	u, err := h.userStore.GetByID(userIDFromToken(c))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	req := newUserUpdateRequest()
	req.populate(u)
	if err := req.bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := h.userStore.Update(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, newUserResponse(u))
}

func userIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
