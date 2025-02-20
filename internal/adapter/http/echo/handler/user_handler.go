package handler

import (
	"github.com/devmyong/todo/backend/backend/internal/adapter/http/echo/dto"
	"github.com/devmyong/todo/backend/backend/internal/domain/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	svc *user.RegisterService
}

func NewUserHandler(svc *user.RegisterService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) RegisterRoutes(e *echo.Echo) {
	e.POST("/user/register", h.registerLocal)
}

func (h *UserHandler) registerLocal(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	u, err := h.svc.RegisterLocal(req.Email, req.Password, req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resp := dto.CreateUserResponse{UserID: u.ID.Hex()}
	return c.JSON(http.StatusCreated, resp)
}
