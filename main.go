package main

import (
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

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	userDlv.New(e, uService)

	e.Logger.Fatal(e.Start(":8000"))
}
