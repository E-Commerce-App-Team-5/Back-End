package delivery

import (
	"ecommerce/config"
	"ecommerce/features/user/domain"
	"ecommerce/middlewares"
	"ecommerce/utils/helper"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type userHandler struct {
	srv domain.Service
}

func New(e *echo.Echo, srv domain.Service) {
	handler := userHandler{srv: srv}
	e.POST("/login", handler.Login())                                                   // LOGIN
	e.POST("/users", handler.Register())                                                // REGISTER
	e.GET("/users", handler.GetUser())                                                  // GET USER & PRODUCT
	e.DELETE("/users", handler.DeleteUser(), middleware.JWT([]byte(config.JWT_SECRET))) // DELETE USER
	e.PUT("/users", handler.UpdateUser(), middleware.JWT([]byte(config.JWT_SECRET)))    // UPDATE USER
}

func (us *userHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("cannot bind data")))
		}

		cnv := ToDomain(input)
		res, token, err := us.srv.Login(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("login failed"))
		}
		res.Token = token
		return c.JSON(http.StatusOK, SuccessResponseWithData("login successful", ToResponse(res, "login")))
	}
}

func (us *userHandler) UpdateUser() echo.HandlerFunc {
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
			input.UserPicture = res
		}

		id := middlewares.ExtractToken(c)
		input.ID = uint(id)
		cnv := ToDomain(input)
		_, err = us.srv.UpdateUser(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}
		return c.JSON(http.StatusCreated, SuccessResponseNoData("update user successful"))
	}
}

func (us *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := middlewares.ExtractToken(c)
		toUint := uint(id)
		_, err := us.srv.DeleteUser(toUint)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}
		return c.JSON(http.StatusOK, SuccessResponseNoData("delete user successful"))
	}
}

func (us *userHandler) GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.QueryParam("username")
		resUser, resProduct, err := us.srv.GetUser(username)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse(err))
		}
		return c.JSON(http.StatusOK, SuccessResponseWithData("Success show all data", ToResponseGetUser(resUser, resProduct)))
	}
}

func (us *userHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input RegisterFormat
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, FailResponse(errors.New("An invalid client request.")))
		}
		cnv := ToDomain(input)
		_, err := us.srv.Register(cnv)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailResponse("username has taken"))
		}

		return c.JSON(http.StatusCreated, SuccessResponseNoData("success register user"))
	}

}
