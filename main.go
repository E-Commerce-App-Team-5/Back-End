package main

import (
	cartDlv "ecommerce/features/cart/delivery"
	cartRepo "ecommerce/features/cart/repository"
	cartSrv "ecommerce/features/cart/services"
	productDlv "ecommerce/features/products/delivery"
	productRepo "ecommerce/features/products/repository"
	productSrv "ecommerce/features/products/services"
	userDlv "ecommerce/features/user/delivery"
	userRepo "ecommerce/features/user/repository"
	userSrv "ecommerce/features/user/services"

	"ecommerce/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	db := database.InitDB()
	uRepo := userRepo.New(db)
	uService := userSrv.New(uRepo)
	pRepo := productRepo.New(db)
	pService := productSrv.New(pRepo)
	cRepo := cartRepo.New(db)
	cService := cartSrv.New(cRepo)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userDlv.New(e, uService)
	productDlv.New(e, pService)
	cartDlv.New(e, cService)

	e.Logger.Fatal(e.Start(":8000"))
}
