package main

import (
	"be_project3team3/config"
	dUser "be_project3team3/feature/user/delivery"
	rUser "be_project3team3/feature/user/repository"
	sUser "be_project3team3/feature/user/services"

	dCart "be_project3team3/feature/cart/delivery"
	rCart "be_project3team3/feature/cart/repository"
	sCart "be_project3team3/feature/cart/services"
	"be_project3team3/utils/database"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	//pemanggilan config
	cfg := config.NewConfig()
	db := database.InitDB(cfg)

	//User
	repoUser := rUser.New(db)
	serUser := sUser.New(repoUser)

	//Cart
	repoCart := rCart.New(db)
	serCart := sCart.New(repoCart)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	dUser.New(e, serUser)
	dCart.New(e, serCart)

	log.Fatal(e.Start(":8000"))

}
