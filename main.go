package main

import (
	cartDlv "ecommerce/features/cart/delivery"
	cartRepo "ecommerce/features/cart/repository"
	cartSrv "ecommerce/features/cart/services"
	chckDlv "ecommerce/features/checkout/delivery"
	chckRepo "ecommerce/features/checkout/repository"
	chckSrv "ecommerce/features/checkout/services"
	productDlv "ecommerce/features/products/delivery"
	productRepo "ecommerce/features/products/repository"
	productSrv "ecommerce/features/products/services"
	userDlv "ecommerce/features/user/delivery"
	userRepo "ecommerce/features/user/repository"
	userSrv "ecommerce/features/user/services"
	historyRepo  "ecommerce/features/history_order/repository"
	historySrv  "ecommerce/features/history_order/services"
	historyDlv  "ecommerce/features/history_order/delivery"

	"ecommerce/utils/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	db := database.InitDB()
	// USER 
	uRepo := userRepo.New(db)
	uService := userSrv.New(uRepo)
	// PRODUCT
	pRepo := productRepo.New(db)
	pService := productSrv.New(pRepo)
	// CART
	cRepo := cartRepo.New(db)
	cService := cartSrv.New(cRepo)
	// CHECKOUT
	ckRepo := chckRepo.New(db)
	ckService := chckSrv.New(ckRepo)
	// HISTORY
	historyRepo := historyRepo.New(db)
	historySrv := historySrv.New(historyRepo)


	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userDlv.New(e, uService)
	productDlv.New(e, pService)
	cartDlv.New(e, cService)
	chckDlv.New(e, ckService)
	historyDlv.New(e, historySrv)

	e.Logger.Fatal(e.Start(":8000"))
}
