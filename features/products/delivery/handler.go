package delivery

import (
	"ecommerce/config"
	"ecommerce/features/products/domain"
	"ecommerce/middlewares"
	"ecommerce/utils/helper"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type productHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := productHandler{srv: srv}
	e.POST("/products", handler.AddProduct(), middleware.JWT([]byte(config.JWT_SECRET)))          // TAMBAH PRODUCT
	e.GET("/products", handler.GetProduct())                                                      // GET PRODUCT
	e.DELETE("/products/:id", handler.DeleteProduct(), middleware.JWT([]byte(config.JWT_SECRET))) // DELETE PRODUCT
	e.PUT("/products", handler.UpdateProduct(), middleware.JWT([]byte(config.JWT_SECRET)))        // UPDATE PRODUCT
}

func (ps *productHandler) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("cannot bind data")))
		}

		file, err := c.FormFile("user_picture")
		if file != nil {
			res, err := helper.UploadProfile(c)
			if err != nil {
				return err
			}
			log.Print(res)
			input.ProductPicture = res
		}

		id := middlewares.ExtractToken(c)
		input.IdUser = uint(id)
		cnv := ToDomain(input)
		res, err := ps.srv.UpdateProduct(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("Success update user", ToResponse(res, "update")))
	}
}

func (ps *productHandler) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return errors.New("cannot convert id")
		}
		_, err = ps.srv.DeleteProduct(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request."))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("Success delete data."))
	}
}

func (ps *productHandler) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		page, err := strconv.Atoi(c.QueryParam("page"))
		if err != nil {
			page = 0
		}
		res, err := ps.srv.GetProduct(page)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("An invalid client request"))
		}
		return c.JSON(http.StatusOK, SuccessResponse("Success show all data", ToResponseProduct(res, "sukses")))
	}
}

func (ps *productHandler) AddProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("an invalid client request")))
		}
		input.IdUser = uint(middlewares.ExtractToken(c))
		file, _ := c.FormFile("product_picture")
		if file != nil {
			res, err := helper.UploadProfile(c)
			if err != nil {
				return err
			}
			log.Print(res)
			input.ProductPicture = res
		}
		cnv := ToDomain(input)
		res, err := ps.srv.AddProduct(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("There is problem on server."))
		}

		return c.JSON(http.StatusCreated, SuccessResponse("success add product", ToResponse(res, "register")))
	}

}
