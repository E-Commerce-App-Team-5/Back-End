package delivery

import (
	"ecommerce/config"
	"ecommerce/features/checkout/domain"
	"ecommerce/middlewares"
	"ecommerce/utils/helper"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type checkoutHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := checkoutHandler{srv: srv}
	e.POST("/checkout", handler.AddCheckout(), middleware.JWT([]byte(config.JWT_SECRET))) // TAMBAH CHECKOUT
	e.POST("/midtrans", handler.UpdateCheckout())
	e.GET("/checkout", handler.GetCheckout(), middleware.JWT([]byte(config.JWT_SECRET)))           // GET CHECKOUT
	e.DELETE("/checkout/:id", handler.DeleteCheckout(), middleware.JWT([]byte(config.JWT_SECRET))) // DELETE CHECKOUT
}

func (cs *checkoutHandler) AddCheckout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input []RegisterFormat
		var inputChck CheckoutFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse("an invalid client request"))
		}

		inputChck.IdPembeli = uint(middlewares.ExtractToken(c))
		cnv, cnvC := ToDomainHistory(input, inputChck)
		cnvC.OrderId = "ORDER-" + (time.Now().Format("02 Jan 06 15:04")) + fmt.Sprintf("%f", rand.Float64())
		cnvC.Status = "pending"
		inputMidtrans := helper.OrderMidtrans(cnvC.OrderId, int64(cnvC.GrossAmount))
		cnvC = ToDomainMidtrans(inputMidtrans, cnvC)
		log.Println(cnvC)
		res, err := cs.srv.AddCheckout(cnv, cnvC)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("There is problem on server."))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add product", ToResponse(res, "register")))
	}
}

func (cs *checkoutHandler) DeleteCheckout() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		_, err = cs.srv.DeleteCheckout(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err.Error()))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("Success delete data."))
	}
}

func (cs *checkoutHandler) GetCheckout() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := uint(middlewares.ExtractToken(c))
		res, err := cs.srv.GetCheckout(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success show all data", ToResponse(res, "get")))
	}
}

func (cs *checkoutHandler) UpdateCheckout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}
		inputMidtrans := helper.CheckMidtrans(input.OrderID)
		res := ToDomainCheckMidtrans(inputMidtrans)
		cs.srv.UpdateCheckout(res)
		return c.JSON(http.StatusOK, SuccessResponseNoData("Success update data."))
	}
}
