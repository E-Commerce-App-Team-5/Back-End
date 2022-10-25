package delivery

import (
	"ecommerce/config"
	"ecommerce/features/checkout/domain"
	"ecommerce/middlewares"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type checkoutHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := checkoutHandler{srv: srv}
	e.POST("/checkout", handler.AddCheckout(), middleware.JWT([]byte(config.JWT_SECRET))) // TAMBAH CART
}

func (cs *checkoutHandler) AddCheckout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input []RegisterFormat
		var inputChck CheckoutFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}
		inputChck.IdPembeli = uint(middlewares.ExtractToken(c))
		cnv := ToDomainHistory(input)
		cnvC := ToDomain(inputChck)
		res, err := cs.srv.AddCheckout(cnv, cnvC)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("There is problem on server."))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add product", ToResponse(res, "register")))
	}
}
