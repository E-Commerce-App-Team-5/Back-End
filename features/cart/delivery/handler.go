package delivery

import (
	"ecommerce/config"
	"ecommerce/features/cart/domain"
	"ecommerce/middlewares"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type cartHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := cartHandler{srv: srv}
	e.POST("/cart", handler.AddCart(), middleware.JWT([]byte(config.JWT_SECRET)))          // TAMBAH CART
	e.GET("/cart", handler.GetCart(), middleware.JWT([]byte(config.JWT_SECRET)))           // GET CART
	e.DELETE("/cart/:id", handler.DeleteCart(), middleware.JWT([]byte(config.JWT_SECRET))) // DELETE CART
	e.PUT("/cart/:id", handler.UpdateCart(), middleware.JWT([]byte(config.JWT_SECRET)))
}

func (cs *cartHandler) UpdateCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		id, err := strconv.Atoi(c.Param("id"))
		input.ID = uint(id)
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("cannot bind data")))
		}

		idUser := middlewares.ExtractToken(c)
		input.IdUser = uint(idUser)
		cnv := ToDomain(input)
		res, err := cs.srv.UpdateCart(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success update cart", ToResponse(res, "update")))
	}
}

func (cs *cartHandler) DeleteCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errors.New("cannot convert id")
		}
		_, err = cs.srv.DeleteCart(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request."))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("Success delete data."))
	}
}

func (cs *cartHandler) GetCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		res, err := cs.srv.GetCart(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success show all data", ToResponseProduct(res, "sukses")))
	}
}

func (cs *cartHandler) AddCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}
		input.IdUser = uint(middlewares.ExtractToken(c))
		cnv := ToDomain(input)
		res, err := cs.srv.AddCart(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add cart", ToResponse(res, "register")))
	}
}
