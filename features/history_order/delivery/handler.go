package delivery

import (
	"ecommerce/config"
	"ecommerce/features/history_order/domain"
	"ecommerce/middlewares"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type historyHandler struct {
	srv domain.Services
}

func New(e *echo.Echo, srv domain.Services) {
	handler := historyHandler{srv: srv}
	e.GET("/historybuy", handler.GetBuy(), middleware.JWT([]byte(config.JWT_SECRET)))
	e.GET("/historysell", handler.GetSell(), middleware.JWT([]byte(config.JWT_SECRET)))
}

func (hh *historyHandler) GetBuy() echo.HandlerFunc {
	return func (c echo.Context) error {
		id := middlewares.ExtractToken(c)
		res, err := hh.srv.GetBuy(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get history buy", ToResponse("buy", res)))
	}
}

func (hh *historyHandler) GetSell() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		res, err := hh.srv.GetSell(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success get history sell", ToResponse("sell", res)))
	}
}

