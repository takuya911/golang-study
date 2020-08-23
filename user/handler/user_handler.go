package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/takuya911/golang-study/conf"
	"github.com/takuya911/golang-study/user/domain"
)

type userHandler struct {
	usecase domain.UserUsecase
}

// NewUserHandler function
func NewUserHandler(e *echo.Echo, u domain.UserUsecase) {
	handler := &userHandler{
		usecase: u,
	}
	e.GET("/user/:user_id", handler.GetUserByID)
	e.POST("/user", handler.StoreUser)
	e.POST("/user/:user_id", handler.UpdateUser)
	e.DELETE("/user/:user_id", handler.DeleteUser)
}

func (h *userHandler) GetUserByID(e echo.Context) error {
	userID, _ := strconv.Atoi(e.Param("user_id"))
	if userID < 1 {
		return e.JSON(http.StatusNotFound, conf.ErrNotFound.Error())
	}

	etx := e.Request().Context()
	user, err := h.usecase.GetByID(etx, userID)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, user)
}

func (h *userHandler) StoreUser(e echo.Context) error {
	etx := e.Request().Context()

	var user domain.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err := h.usecase.Store(etx, &user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, user)
}

func (h *userHandler) UpdateUser(e echo.Context) error {
	userID, err := strconv.Atoi(e.Param("user_id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, conf.ErrNotFound.Error())
	}
	etx := e.Request().Context()

	var user domain.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusUnprocessableEntity, err.Error())
	}
	user.ID = userID

	err = h.usecase.Update(etx, &user)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusCreated, user)
}

func (h *userHandler) DeleteUser(e echo.Context) error {
	userID, err := strconv.Atoi(e.Param("user_id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, conf.ErrNotFound.Error())
	}

	etx := e.Request().Context()
	if err := h.usecase.Delete(etx, userID); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.NoContent(http.StatusNoContent)
}
