package main

import (
	"be_project3team3/config"
	dOrder "be_project3team3/feature/order/delivery"
	rOrder "be_project3team3/feature/order/repository"
	sOrder "be_project3team3/feature/order/services"
	dProduct "be_project3team3/feature/product/delivery"
	rProduct "be_project3team3/feature/product/repository"
	sProduct "be_project3team3/feature/product/services"
	dUser "be_project3team3/feature/user/delivery"
	rUser "be_project3team3/feature/user/repository"
	sUser "be_project3team3/feature/user/services"
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

	mdlUser := rUser.New(db)
	mdlProduct := rProduct.New(db)
	mdlOrder := rOrder.New(db)

	serUser := sUser.New(mdlUser)
	serProduct := sProduct.New(mdlProduct)
	serOrder := sOrder.New(mdlOrder)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	dUser.New(e, serUser)
	dProduct.New(e, serProduct)
	dOrder.New(e, serOrder)

	log.Fatal(e.Start(":8000"))

}
